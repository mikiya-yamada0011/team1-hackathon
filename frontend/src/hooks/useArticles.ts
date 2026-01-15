'use client';

import { useQuery } from '@tanstack/react-query';
import { getGetApiArticlesQueryOptions } from '@/generated/api/記事-articles/記事-articles';
import type { GetApiArticlesParams } from '@/generated/models';

export const useArticles = (params?: GetApiArticlesParams) => {
  return useQuery(getGetApiArticlesQueryOptions(params));
};
