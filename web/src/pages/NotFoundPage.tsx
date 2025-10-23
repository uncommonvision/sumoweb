export default function NotFoundPage() {
  return (
    <div className="flex flex-col items-center justify-center min-h-[60vh] space-y-4">
      <h1 className="text-6xl font-bold text-foreground">404</h1>
      <h2 className="text-2xl font-semibold text-foreground">Page Not Found</h2>
      <p className="text-lg text-muted-foreground">
        The page you are looking for does not exist.
      </p>
    </div>
  )
}
