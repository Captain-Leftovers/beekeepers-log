/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go", ],
  theme: {
    extend: {
      colors: {
// add primary, primary-foreground, background, background-foreground, and accent colors
        primary: {
          DEFAULT: '#0D47A1',
          foreground: '#FFFFFF',
        },
        background: {
          DEFAULT: '#FFFFFF',
          foreground: '#000000',
        },
        accent: {
          DEFAULT: '#FF6F00',
          foreground: '#000000',
        },
      },
    },
  },
  plugins: [],
}

