'use client';

import { ThemeProvider as NextThemesProvider } from 'next-themes';
import { NextUIProvider } from '@nextui-org/react';
import { GoogleOAuthProvider } from '@react-oauth/google';

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <NextThemesProvider defaultTheme='system' attribute='class'>
      <NextUIProvider>
        <GoogleOAuthProvider clientId={process.env.GOOGLE_CLIENT_ID ?? ''}>
          {children}
        </GoogleOAuthProvider>
      </NextUIProvider>
    </NextThemesProvider>
  );
}
