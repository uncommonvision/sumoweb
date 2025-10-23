import type { ReactNode } from 'react'
import NoAuthHeader from './NoAuthHeader'
import { KeyboardShortcutsOverlay } from '@/components/ui'

interface NoAuthLayoutProps {
  className: string,
  children: ReactNode,
}

export default function NoAuthLayout({ className, children }: NoAuthLayoutProps) {
  return (
    <div className={`min-h-screen bg-background font-sans antialiased ${className}`}>
      <NoAuthHeader />
      <main className="container w-max-3xl mx-auto px-4 pt-6 pb-12">
        {children}
      </main>
      <KeyboardShortcutsOverlay />
    </div>
  )
}
