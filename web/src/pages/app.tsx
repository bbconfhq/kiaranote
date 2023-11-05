import {
  DndProvider,
  getBackendOptions,
  MultiBackend,
  Tree,
  NodeModel,
  getDescendants,
  DropOptions
} from '@minoru/react-dnd-treeview';
import { Box, Section } from '@radix-ui/themes';
import { AxiosError, HttpStatusCode } from 'axios';
import React, { useEffect, useState } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';

import { css } from '../../styled-system/css';
import { api } from '../api';
import Node from '../components/document-tree-node';
import useTreeOpenHandler from '../hooks/useTreeOpenHandler';

type NoteItem = {
  id: number;
  user_id: number;
  title: string;
  content: string;
  is_private: boolean;
  child_notes: NoteItem[] | null;
  create_dt: string;
  update_dt: string;
};

const treeRootStyle = css({
  listStyleType: 'none',
  paddingInlineStart: 0,
  padding: 10,
  position: 'relative',
});

const placeholderStyle = css({
  position: 'relative',
});

const Placeholder = ({ depth }: { depth: number }) => {
  return (
    <div
      style={{
        position: 'absolute',
        top: 0,
        right: 0,
        height: 4,
        left: depth * 24,
        transform: 'translateY(-50%)',
        backgroundColor: '#81a9e0',
        zIndex: 100
      }}
    />
  );
};

type TreeNodeDataState = {
  nodes: NodeModel<NoteItem>[];
  childrenMap: Map<number, boolean>;
};

const resolveChildrenMap = (nodes: NodeModel<NoteItem>[]) => {
  const map = new Map<number, boolean>();
  nodes.forEach((node) => {
    map.set(node.parent as number, true);
  });
  return map;
};

const reorderArray = (array: NodeModel<NoteItem>[], sourceIndex: number, targetIndex: number) => {
  const newArray = [...array];
  const element = newArray.splice(sourceIndex, 1)[0];
  newArray.splice(targetIndex, 0, element);
  return newArray;
};


const MainPage = () => {
  const navigate = useNavigate();
  const [treeData, setTreeData] = useState<TreeNodeDataState>({
    nodes: [],
    childrenMap: new Map(),
  });
  const { ref, toggle } = useTreeOpenHandler();

  const handleDrop = (_: unknown, e: DropOptions) => {
    const { dragSourceId, dropTargetId, destinationIndex } = e;
    if (dragSourceId == null || dropTargetId == null) {
      return;
    }
    const start = treeData.nodes.find((v) => v.id === dragSourceId);
    const end = treeData.nodes.find((v) => v.id === dropTargetId);

    if (
      start?.parent === dropTargetId &&
      start &&
      typeof destinationIndex === 'number'
    ) {
      setTreeData((prev) => {
        const nodes = reorderArray(
          prev.nodes,
          prev.nodes.indexOf(start),
          destinationIndex
        );
        return {
          nodes,
          childrenMap: resolveChildrenMap(nodes),
        };
      });
    }

    if (
      start?.parent !== dropTargetId &&
      start &&
      typeof destinationIndex === 'number'
    ) {
      if (
        getDescendants(treeData.nodes, dragSourceId).find(
          (el) => el.id === dropTargetId
        ) ||
        dropTargetId === dragSourceId ||
        (end && !end?.droppable)
      )
        return;
      setTreeData((prev) => {
        const nodes = reorderArray(
          prev.nodes,
          prev.nodes.indexOf(start),
          destinationIndex
        );
        const movedElement = nodes.find((el) => el.id === dragSourceId);
        if (movedElement) {
          movedElement.parent = dropTargetId;
        }
        return {
          nodes,
          childrenMap: resolveChildrenMap(nodes),
        };
      });
    }
  };

  const getTreeNodeModel = (note: NoteItem, rootId: number): NodeModel<NoteItem>[] => {
    const model = {
      id: note.id,
      parent: rootId,
      text: note.title,
      droppable: true,
      data: note,
    };

    return [
      model,
      ...note.child_notes?.flatMap((item) => getTreeNodeModel(item, note.id)) ?? [],
    ];
  };
  useEffect(() => {
    (async function() {
      try {
        const response = await api.getNotes();
        setTreeData(() => {
          const nodes = [...response.data.data.public, ...response.data.data.private].flatMap((item) => getTreeNodeModel(item, 0));
          return {
            nodes,
            childrenMap: resolveChildrenMap(nodes),
          };
        });
      } catch (err) {
        if (err instanceof AxiosError) {
          if (err.response?.status === HttpStatusCode.Unauthorized) {
            return navigate('/sign-in');
          }
          console.error(err);
        }
      }
    })();
  }, []);

  return (
    <Box height={'100%'}>
      <Section>
        <DndProvider backend={MultiBackend} options={getBackendOptions()}>
          <Tree tree={treeData.nodes}
            ref={ref}
            classes={{
              root: treeRootStyle,
              placeholder: placeholderStyle,
            }}
            sort={false}
            rootId={0}
            canDrop={() => true}
            onDrop={handleDrop}
            dropTargetOffset={5}
            placeholderRender={(node, { depth }) => (
              <Placeholder depth={depth} />
            )}
            render={(node, { depth, isOpen }) => (
              <Node
                node={node}
                isLeaf={!treeData.childrenMap.get(node.id as number)}
                depth={depth}
                isOpen={isOpen}
                onClick={() => {
                  if (node.droppable) {
                    toggle(node.id);
                  }
                }}
              />
            )} />
        </DndProvider>
      </Section>
      <Section>
        <Outlet />
      </Section>
    </Box>
  );
};

export default MainPage;
