import { defineConfig } from 'orval';

export default defineConfig({
  blog: {
    input: {
      target: './src/generated/openapi.json',
    },
    output: {
      mode: 'tags-split',
      target: './src/generated/api',
      schemas: './src/generated/models',
      client: 'react-query',
      mock: false,
      override: {
        mutator: {
          path: './src/lib/api-client.ts',
          name: 'customInstance',
        },
        query: {
          useQuery: true,
          useInfinite: false,
          signal: true,
        },
      },
    },
  },
});
