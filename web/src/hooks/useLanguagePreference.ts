import { useState, useEffect } from 'react'
import type { Language } from '@/types'

export function useLanguagePreference(): [Language, (lang: Language) => void] {
  const [language, setLanguage] = useState<Language>(() => {
    const saved = localStorage.getItem('language')
    return (saved === 'en' || saved === 'jp') ? saved : 'en'
  })

  useEffect(() => {
    localStorage.setItem('language', language)
  }, [language])

  return [language, setLanguage]
}
