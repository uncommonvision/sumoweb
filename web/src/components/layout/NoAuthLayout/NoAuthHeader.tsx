import ThemeToggle from '../../ui/ThemeToggle'

export default function NoAuthHeader() {
  return (
    <header className="sticky top-0 z-50 w-full border-b border-border/40 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container mx-auto flex h-16 items-center justify-between px-2 sm:px-4">
        <div className="flex items-center space-x-4">
          <div className="flex items-center space-x-2">
            <img src="/logo.png" className="w-[32px] h-[32px]" />
            <h1 className="text-xl font-bold text-foreground">Sumo</h1>
          </div>
        </div>

        <div className="flex items-center space-x-1 sm:space-x-2">
          <ThemeToggle />
        </div>
      </div>
    </header>
  )
}
