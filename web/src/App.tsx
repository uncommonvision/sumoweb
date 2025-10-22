import { Outlet } from 'react-router-dom'
import { ThemeProvider } from "@/components/theme-provider"
import { LanguageProvider } from "@/contexts/LanguageContext"

function App() {
  return (
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <LanguageProvider>
        <Outlet />
      </LanguageProvider>
    </ThemeProvider>
  )
}

export default App
