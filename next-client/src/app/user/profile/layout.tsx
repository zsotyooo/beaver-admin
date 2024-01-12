import { SidebarLayout } from '@/modules/common/components/layouts/sidebar-layout';

export default function Layout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <SidebarLayout>
      <div className='mx-auto my-14 flex w-full max-w-[95rem] flex-col gap-4'>
        {children}
      </div>
    </SidebarLayout>
  );
}
