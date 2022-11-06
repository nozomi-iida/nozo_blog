export const pagesPath = {
  "articles": {
    _id: (id: string | number) => ({
      $url: (url?: { hash?: string }) => ({ pathname: '/articles/[id]' as const, query: { id }, hash: url?.hash })
    })
  },
  $url: (url?: { hash?: string }) => ({ pathname: '/' as const, hash: url?.hash })
}

export type PagesPath = typeof pagesPath

export const staticPath = {
  blog_test_jpeg: '/blog_test.jpeg',
  favicon_ico: '/favicon.ico',
  vercel_svg: '/vercel.svg'
} as const

export type StaticPath = typeof staticPath