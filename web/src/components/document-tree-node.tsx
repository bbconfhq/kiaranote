import { NodeModel } from '@minoru/react-dnd-treeview';
import * as ContextMenu from '@radix-ui/react-context-menu';
import { FileTextIcon } from '@radix-ui/react-icons';
import React, { MouseEventHandler } from 'react';

import { css } from '../../styled-system/css';
import { styled } from '../../styled-system/jsx';
import { api } from '../api';

const TREE_X_OFFSET = 22;

const Wrapper = styled('div', {
  base: {
    alignItems: 'center',
    display: 'grid',
    gridTemplateColumns: 'auto auto 1fr',
    height: '32px',
    paddingInlineEnd: '8px',
    paddingInlineStart: '8px',
    borderRadius: '4px',
    cursor: 'pointer',
    whiteSpace: 'nowrap',
    position: 'relative',
    zIndex: 3,

    _hover: {
      backgroundColor: 'rgba(0, 0, 0, 0.04)',
    },
  },
});

const ExpandIconWrapper = styled('div', {
  base: {
    alignItems: 'center',
    fontSize: '0',
    cursor: 'pointer',
    display: 'flex',
    height: '24px',
    justifyContent: 'center',
    width: '24px',
    transform: 'rotate(0deg)',
  },
  variants: {
    expanded: {
      true: {
        transform: 'rotate(180deg)',

        '& svg path': {
          fill: '#4f5272',
        }
      }
    }
  }
});

const Label = styled('div', {
  base: {
    paddingInlineStart: '8px',
    overflow: 'hidden',
    textOverflow: 'ellipsis',
  }
});

const ContextMenuWrapper=({
  onDelete,
  children
}: {
  onDelete: (e: Event) => void;
  children: React.ReactNode
})=>{
  return (
    <ContextMenu.Root>
      <ContextMenu.Trigger>
        {children}
      </ContextMenu.Trigger>
      <ContextMenu.Portal>
        <ContextMenu.Content 
          className={css({
            backgroundColor: 'white',
            borderRadius: '8px',
            boxShadow: '0px 4px 16px rgba(0, 0, 0, 0.12)',
            minWidth: '200px',
            padding: '8px',
          })}
        >
          <ContextMenu.Item onSelect={onDelete}>삭제</ContextMenu.Item>
        </ContextMenu.Content>
      </ContextMenu.Portal>
    </ContextMenu.Root>
  );
};

const Node = ({node, depth, isOpen, isLeaf, onClick }: {
  node: NodeModel;
  depth: number;
  isLeaf: boolean;
  isOpen: boolean;
  onClick: (id: string | number) => void;
}) => {
  const indent = depth * TREE_X_OFFSET;

  const handleToggle: MouseEventHandler = (e) => {
    e.stopPropagation();
    onClick(node.id);
  };

  const handleContextMenuDelete = (e: Event) => {
    e.preventDefault();
    e.stopPropagation();
    api.deleteNote(node.id as number);
  };

  return (
    <ContextMenuWrapper onDelete={handleContextMenuDelete}>
      <Wrapper style={{ marginInlineStart: indent }} onClick={handleToggle}>
        <ExpandIconWrapper expanded={isOpen}>
          {!isLeaf && node.droppable && (
            <svg
              width='16'
              height='16'
              viewBox='0 0 16 16'
              fill='none'
              xmlns='http://www.w3.org/2000/svg'
            >
              <path
                d='M10.5866 5.99969L7.99997 8.58632L5.41332 5.99969C5.15332 5.73969 4.73332 5.73969 4.47332 5.99969C4.21332 6.25969 4.21332 6.67965 4.47332 6.93965L7.5333 9.99965C7.59497 10.0615 7.66823 10.1105 7.7489 10.144C7.82957 10.1775 7.91603 10.1947 8.0033 10.1947C8.09063 10.1947 8.1771 10.1775 8.25777 10.144C8.33837 10.1105 8.41163 10.0615 8.4733 9.99965L11.5333 6.93965C11.7933 6.67965 11.7933 6.25969 11.5333 5.99969C11.2733 5.74635 10.8466 5.73969 10.5866 5.99969Z'
                fill='black'
              />
            </svg>
          )}
        </ExpandIconWrapper>
        <FileTextIcon />
        <Label>{node.text}</Label>
      </Wrapper>
    </ContextMenuWrapper>

  );
};

export default Node;
