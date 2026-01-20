'use client';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useEffect, useState } from 'react';
import { useAuth } from '@/hooks/useAuth';

export function Providers({ children }: { children: React.ReactNode }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 60 * 1000, // 1分
            refetchOnWindowFocus: false,
          },
        },
      }),
  );

  return (
    <QueryClientProvider client={queryClient}>
      <AuthInitializer>{children}</AuthInitializer>
    </QueryClientProvider>
  );
}

function AuthInitializer({ children }: { children: React.ReactNode }) {
  const initialize = useAuth((state) => state.initialize);
  const [isInitialized, setIsInitialized] = useState(false);

  useEffect(() => {
    initialize().finally(() => setIsInitialized(true));
  }, [initialize]);

  if (!isInitialized) {
    return null; // または <div>Loading...</div>
  }

  return <>{children}</>;
}
