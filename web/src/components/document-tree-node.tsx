import { NodeModel } from '@minoru/react-dnd-treeview';
import React, { MouseEventHandler } from 'react';

import { styled } from '../../styled-system/jsx';

const TREE_X_OFFSET = 22;

const NodeIcon = () => {
  return (
    <svg
      width='16'
      height='16'
      viewBox='0 0 16 16'
      fill='none'
      xmlns='http://www.w3.org/2000/svg'
    >
      <path
        d='M3 2.5C3 2.22386 3.22386 2 3.5 2H9.08579C9.21839 2 9.34557 2.05268 9.43934 2.14645L11.8536 4.56066C11.9473 4.65443 12 4.78161 12 4.91421V12.5C12 12.7761 11.7761 13 11.5 13H3.5C3.22386 13 3 12.7761 3 12.5V2.5ZM3.5 1C2.67157 1 2 1.67157 2 2.5V12.5C2 13.3284 2.67157 14 3.5 14H11.5C12.3284 14 13 13.3284 13 12.5V4.91421C13 4.51639 12.842 4.13486 12.5607 3.85355L10.1464 1.43934C9.86514 1.15804 9.48361 1 9.08579 1H3.5ZM4.5 4C4.22386 4 4 4.22386 4 4.5C4 4.77614 4.22386 5 4.5 5H7.5C7.77614 5 8 4.77614 8 4.5C8 4.22386 7.77614 4 7.5 4H4.5ZM4.5 7C4.22386 7 4 7.22386 4 7.5C4 7.77614 4.22386 8 4.5 8H10.5C10.7761 8 11 7.77614 11 7.5C11 7.22386 10.7761 7 10.5 7H4.5ZM4.5 10C4.22386 10 4 10.2239 4 10.5C4 10.7761 4.22386 11 4.5 11H10.5C10.7761 11 11 10.7761 11 10.5C11 10.2239 10.7761 10 10.5 10H4.5Z'
        fill='currentColor'
        fillRule='evenodd'
        clipRule='evenodd'
      >
      </path>
    </svg>
  );
};

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

const Node = ({node, depth, isOpen, onClick }: {
  node: NodeModel;
  depth: number;
  isOpen: boolean;
  onClick: (id: string | number) => void;
}) => {
  const indent = depth * TREE_X_OFFSET;

  const handleToggle: MouseEventHandler = (e) => {
    e.stopPropagation();
    onClick(node.id);
  };

  return (
    <Wrapper style={{ marginInlineStart: indent }} onClick={handleToggle}>
      <ExpandIconWrapper expanded={isOpen}>
        {node.droppable && (
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
      <NodeIcon />
      <Label>{node.text}</Label>
    </Wrapper>
  );
};

export default Node;
