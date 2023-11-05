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

type NoteResponse = {
  public: NoteItem[];
  private: NoteItem[];
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

const reorderArray = (array: NodeModel<NoteItem>[], sourceIndex: number, targetIndex: number) => {
  const newArray = [...array];
  const element = newArray.splice(sourceIndex, 1)[0];
  newArray.splice(targetIndex, 0, element);
  return newArray;
};


const MainPage = () => {
  const navigate = useNavigate();
  const [notes, setNotes] = useState<NoteResponse | null>(null);
  const [treeData, setTreeData] = useState<NodeModel<NoteItem>[]>([]);
  const { ref, toggle } = useTreeOpenHandler();

  const handleDrop = (_: unknown, e: DropOptions) => {
    const { dragSourceId, dropTargetId, destinationIndex } = e;
    if (dragSourceId == null || dropTargetId == null) {
      return;
    }
    const start = treeData.find((v) => v.id === dragSourceId);
    const end = treeData.find((v) => v.id === dropTargetId);

    if (
      start?.parent === dropTargetId &&
      start &&
      typeof destinationIndex === 'number'
    ) {
      setTreeData((prev) => {
        const output = reorderArray(
          prev,
          prev.indexOf(start),
          destinationIndex
        );
        return output;
      });
    }

    if (
      start?.parent !== dropTargetId &&
      start &&
      typeof destinationIndex === 'number'
    ) {
      if (
        getDescendants(treeData, dragSourceId).find(
          (el) => el.id === dropTargetId
        ) ||
        dropTargetId === dragSourceId ||
        (end && !end?.droppable)
      )
        return;
      setTreeData((prev) => {
        const output = reorderArray(
          prev,
          prev.indexOf(start),
          destinationIndex
        );
        const movedElement = output.find((el) => el.id === dragSourceId);
        if (movedElement) movedElement.parent = dropTargetId;
        return output;
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
        setNotes(response.data.data);
        setTreeData([...response.data.data.public, ...response.data.data.private].flatMap((item) => getTreeNodeModel(item, 0)));
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

  if (notes == null) {
    return null;
  }

  return (
    <Box height={'100%'}>
      <Section>
        <DndProvider backend={MultiBackend} options={getBackendOptions()}>
          <Tree tree={treeData}
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
            render={(node, { depth, isOpen, isDropTarget }) => (
              <Node
                node={node}
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
