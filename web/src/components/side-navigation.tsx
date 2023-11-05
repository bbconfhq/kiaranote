import { Text, Heading, Box, Flex } from '@radix-ui/themes';
import React from 'react';
import { Link } from 'react-router-dom';

import styles from './side-navigation.module.css';

interface DocsNavProps {
  routes: {
    label?: string;
    pages: {
      title: string;
      slug: string;
      icon?: React.ReactNode;
    }[];
  }[];
}

export const SideNavigation = ({ routes }: DocsNavProps) => {
  return (
    <Box>
      {routes.map((section, i) => (
        <Box key={section.label ?? i} mb='4'>
          {section.label && (
            <Box py='2' px='3'>
              <Heading as='h4' size={{ initial: '3', md: '2' }}>
                {section.label}
              </Heading>
            </Box>
          )}

          {section.pages.map((page) => (
            <NavigationItem key={page.slug} href={page.slug} active={false}>
              <Flex gap='2' align='center'>
                {page.icon}
                <Text size={{ initial: '3', md: '2' }}>{page.title}</Text>
              </Flex>
            </NavigationItem>
          ))}
        </Box>
      ))}
    </Box>
  );
};

interface NavigationItemProps {
  children: React.ReactNode;
  active?: boolean;
  href: string;
  className?: string;
}

const classNames = (...names: Array<unknown>) => {
  return names.filter(Boolean).join(' ');
};

const NavigationItem = ({ active, href, ...props }: NavigationItemProps) => {
  const className = classNames(styles.DocsNavItem, active && styles.active);
  const ref = React.useRef<HTMLAnchorElement>(null);

  return (
    <Link to={`/${href}`} ref={ref} className={className} {...props}>
    </Link>
  );
};
