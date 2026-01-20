'use client';

import { useQuery } from '@tanstack/react-query';
import { getGetApiArticlesSlugQueryOptions } from '@/generated/api/記事-articles/記事-articles';

export const useArticle = (slug: string) => {
  return useQuery(getGetApiArticlesSlugQueryOptions(slug));
};
