import cx from 'classnames';

export default function CenteredLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <main
      className={cx(
        'flex min-h-screen w-full items-center justify-center p-20'
      )}
    >
      <div>{children}</div>
    </main>
  );
}
