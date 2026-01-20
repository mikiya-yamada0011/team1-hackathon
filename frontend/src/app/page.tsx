'use client';

import { Filter, Search } from 'lucide-react';
import Image from 'next/image';
import Link from 'next/link';
import { useState } from 'react';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { Input } from '@/components/ui/input';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import type {
  ArticleResponse,
  GetApiArticlesDepartment,
  GetApiArticlesStatus,
} from '@/generated/models';
import { useArticles } from '@/hooks/useArticles';
import { useAuth } from '@/hooks/useAuth';

const DEPARTMENTS = [
  {
    id: 1,
    label: 'マーケティング',
    value: 'マーケティング' as GetApiArticlesDepartment,
  },
  { id: 2, label: '開発', value: '開発' as GetApiArticlesDepartment },
  { id: 3, label: '組織管理', value: '組織管理' as GetApiArticlesDepartment },
];

const STATUS_OPTIONS = [
  { label: 'すべて（公開 + 内部）', value: 'all' as GetApiArticlesStatus },
  { label: '外部公開のみ', value: 'public' as GetApiArticlesStatus },
  { label: '内部公開のみ', value: 'internal' as GetApiArticlesStatus },
];

export default function Home() {
  const [searchQuery, setSearchQuery] = useState('');
  const [activeTab, setActiveTab] = useState<GetApiArticlesDepartment>(
    '開発' as GetApiArticlesDepartment,
  );
  const [statusFilter, setStatusFilter] = useState<GetApiArticlesStatus>('all');
  const { isAuthenticated } = useAuth();

  // APIから記事を取得
  // ログイン済みの場合はユーザーが選択したステータスで取得
  // ゲストの場合は自動的にpublicのみが取得される
  const { data, isLoading, error } = useArticles({
    department: activeTab,
    page: 1,
    limit: 100,
    status: isAuthenticated ? statusFilter : undefined, // ログイン済みならユーザーの選択、ゲストはundefined
  });

  const filteredArticles = (data?.articles || []).filter((article) => {
    if (!searchQuery) return true;
    const query = searchQuery.toLowerCase();
    return (
      article.title?.toLowerCase().includes(query) ||
      article.author?.name?.toLowerCase().includes(query)
    );
  });

  const formatDate = (dateString?: string) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('ja-JP', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    });
  };

  return (
    <main className="min-h-screen bg-slate-50 text-foreground">
      {/* Header Image Section */}
      <div className="w-full h-32 md:h-64 lg:h-80 relative overflow-hidden bg-slate-200">
        <Image src="/header.png" alt="Header Image" fill className="object-cover" priority />
      </div>

      <div className="container mx-auto px-4 py-4 md:py-12 max-w-[1600px]">
        {/* Info Cards Section */}
        <div className="grid grid-cols-2 gap-2 mb-4 md:mb-6">
          <Link
            href="https://a4-home-page.vercel.app/home"
            target="_blank"
            rel="noopener noreferrer"
            className="group"
          >
            <div className="bg-white border border-slate-200 rounded-lg p-2 md:p-4 hover:border-slate-400 transition-colors">
              <div className="flex items-center gap-2 md:gap-3">
                <div className="w-8 h-8 md:w-10 md:h-10 bg-slate-100 rounded flex items-center justify-center flex-shrink-0">
                  <svg
                    className="w-4 h-4 md:w-5 md:h-5 text-slate-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    aria-label="ホームページアイコン"
                  >
                    <title>ホームページアイコン</title>
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                    />
                  </svg>
                </div>
                <div className="flex-1 min-w-0">
                  <p className="hidden md:block text-[11px] text-slate-500 mb-0.5">
                    神大生に最高の学生生活を。
                  </p>
                  <h3 className="text-[11px] md:text-sm font-semibold text-slate-900">
                    A4 ホームページ
                  </h3>
                </div>
              </div>
            </div>
          </Link>

          <Link
            href="https://a4-home-page.vercel.app/product"
            target="_blank"
            rel="noopener noreferrer"
            className="group"
          >
            <div className="bg-white border border-slate-200 rounded-lg p-2 md:p-4 hover:border-slate-400 transition-colors">
              <div className="flex items-center gap-2 md:gap-3">
                <div className="w-8 h-8 md:w-10 md:h-10 bg-slate-100 rounded flex items-center justify-center flex-shrink-0">
                  <svg
                    className="w-4 h-4 md:w-5 md:h-5 text-slate-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    aria-label="アプリ一覧アイコン"
                  >
                    <title>アプリ一覧アイコン</title>
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z"
                    />
                  </svg>
                </div>
                <div className="flex-1 min-w-0">
                  <p className="hidden md:block text-[11px] text-slate-500 mb-0.5 leading-tight">
                    様々なプロダクトをご利用ください
                  </p>
                  <h3 className="text-[11px] md:text-sm font-semibold text-slate-900">
                    A4のアプリ一覧
                  </h3>
                </div>
              </div>
            </div>
          </Link>
        </div>

        {/* Filter Section */}
        <div className="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
          <div className="w-full md:w-auto" />
          <div className="flex items-center gap-3 w-full md:w-auto">
            {/* ログイン済みの場合のみステータスフィルターを表示 */}
            {isAuthenticated && (
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" className="gap-2 bg-white border-slate-200">
                    <Filter className="h-4 w-4" />
                    {STATUS_OPTIONS.find((opt) => opt.value === statusFilter)?.label}
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end" className="w-56">
                  {STATUS_OPTIONS.map((option) => (
                    <DropdownMenuItem
                      key={option.value}
                      onClick={() => setStatusFilter(option.value)}
                      className={statusFilter === option.value ? 'bg-slate-100' : ''}
                    >
                      {option.label}
                    </DropdownMenuItem>
                  ))}
                </DropdownMenuContent>
              </DropdownMenu>
            )}
            <div className="relative w-full md:w-72">
              <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                type="search"
                placeholder="Search articles..."
                className="pl-9 bg-white border-slate-200"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
              />
            </div>
          </div>
        </div>

        {/* Tabs Section */}
        <Tabs
          defaultValue="開発"
          className="w-full"
          onValueChange={(value) => setActiveTab(value as GetApiArticlesDepartment)}
        >
          <TabsList className="grid w-full grid-cols-3 mb-8 h-12 bg-white border border-slate-200">
            {DEPARTMENTS.map((dept) => (
              <TabsTrigger
                key={dept.id}
                value={dept.value}
                className="text-base font-medium data-[state=active]:bg-primary data-[state=active]:text-primary-foreground transition-all"
              >
                {dept.label}
              </TabsTrigger>
            ))}
          </TabsList>

          {DEPARTMENTS.map((dept) => (
            <TabsContent key={dept.id} value={dept.value} className="mt-0">
              {isLoading ? (
                <div className="col-span-full text-center py-12 text-muted-foreground">
                  Loading...
                </div>
              ) : error ? (
                <div className="col-span-full text-center py-12 text-red-500">
                  Error loading articles: {error.message}
                </div>
              ) : (
                <div className="grid grid-cols-2 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 md:gap-6">
                  {filteredArticles.length > 0 ? (
                    filteredArticles.map((article) => (
                      <ArticleCard key={article.id} article={article} formatDate={formatDate} />
                    ))
                  ) : (
                    <div className="col-span-full text-center py-12 text-muted-foreground">
                      No articles found in {dept.label}.
                    </div>
                  )}
                </div>
              )}
            </TabsContent>
          ))}
        </Tabs>
      </div>
    </main>
  );
}

function ArticleCard({
  article,
  formatDate,
}: {
  article: ArticleResponse;
  formatDate: (date?: string) => string;
}) {
  // 外部記事の場合は外部URLへ、内部記事の場合は詳細ページへ
  const isExternal = article.article_type === 'external' && article.external_url;
  const href = isExternal ? (article.external_url ?? '') : `/detail/${article.slug}`;

  const handleCardClick = () => {
    if (isExternal && article.external_url) {
      window.open(article.external_url, '_blank', 'noopener,noreferrer');
    } else {
      window.location.href = href;
    }
  };

  return (
    <Card
      className="group hover:shadow-xl transition-all duration-300 border-none shadow-sm bg-white flex flex-col h-full overflow-hidden cursor-pointer"
      onClick={handleCardClick}
    >
      {/* Thumbnail Image */}
      <div className="relative w-full aspect-video overflow-hidden bg-slate-100">
        {article.thumbnail_url ? (
          <Image
            src={article.thumbnail_url}
            alt={article.title || 'Article thumbnail'}
            fill
            className="object-cover group-hover:scale-105 transition-transform duration-500"
          />
        ) : (
          <div className="w-full h-full flex items-center justify-center text-muted-foreground">
            No Image
          </div>
        )}
      </div>

      <CardHeader className="pb-2 pt-4 px-5 space-y-0">
        <div className="flex justify-between items-start w-full mb-2">
          <Link
            href={`/users/${article.author?.id}`}
            className="flex items-center gap-2 hover:underline z-10"
            onClick={(e) => e.stopPropagation()}
          >
            <Avatar className="h-6 w-6 border border-slate-100">
              <AvatarImage src={article.author?.icon_url} alt={article.author?.name} />
              <AvatarFallback className="text-[10px]">
                {article.author?.name?.slice(0, 2)}
              </AvatarFallback>
            </Avatar>
            <span className="text-xs font-medium text-slate-600">{article.author?.name}</span>
          </Link>
          <span className="text-xs text-slate-400">{formatDate(article.created_at)}</span>
        </div>
      </CardHeader>

      <CardContent className="flex-grow py-0 px-5">
        <h3 className="text-lg font-bold leading-snug mb-3 group-hover:text-primary transition-colors line-clamp-2 text-slate-800">
          {article.title}
        </h3>
      </CardContent>

      <CardFooter className="pt-2 pb-4 px-5 text-xs text-muted-foreground flex justify-between items-center mt-auto">
        {article.external_url && (
          <span className="text-primary flex items-center gap-1 font-medium text-xs">
            外部リンク
          </span>
        )}
      </CardFooter>
    </Card>
  );
}
