/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./**/*.html', './**/*.templ'],
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
			animation: {
				grow: 'grow 1s ease-out',
			},
			keyframes: {
				grow: {
					//width grows from 0 to 100%
					'0%': { width: '0%', opacity: 0 },
					'80%': { width: '110%' },
					'100%': { width: '100%', opacity: 1 },
				},
			},
		},
	},
	plugins: [],
}
