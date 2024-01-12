import CenteredLayout from '@/modules/common/components/layouts/centered-layout';

export default function Layout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return <CenteredLayout>{children}</CenteredLayout>;
}
