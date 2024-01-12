'use client';
import React from 'react';

import { SidebarContext } from './layout-context';
import { useLockedBody } from '../../hooks/useBodyLock';

interface Props {
  children: React.ReactNode;
}

export const SidebarLayout = ({ children }: Props) => {
  const [sidebarOpen, setSidebarOpen] = React.useState(false);
  const [_, setLocked] = useLockedBody(false);
  const handleToggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
    setLocked(!sidebarOpen);
  };

  return (
    <main className='flex'>
      <section>sidebar</section>
      <section>{children}</section>
    </main>
  );
};
