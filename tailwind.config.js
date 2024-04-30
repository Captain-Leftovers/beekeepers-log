/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./**/*.html', './**/*.templ'],
	theme: {
		extend: {
			colors: {
				// add primary, primary-foreground, background, background-foreground, and accent colors
				primary: {
					DEFAULT: '#0D47A1',
					disabled: '#90CAF9',
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
				danger: {
					DEFAULT: '#D32F2F',
					foreground: '#FFFFFF',
				},
				warning: {
					DEFAULT: '#FFA000',
					foreground: '#000000',
				},
			},
			animation: {
				grow: 'grow 1s ease-out',
				'fade-in': 'fade-in 0.5s ease-out',
				'grow-height': 'grow-height 5s ease-out',
				'fade-out': 'fade-out 0.5s ease-out',
			},
			keyframes: {
				'fade-in': {
					'0%': { opacity: 0 },
					'100%': { opacity: 1 },
				},
				'fade-out': {
					'0%': { opacity: 1 },
					'100%': { opacity: 0 },
				},

				'grow-height': {
					'0%': { height: '0%', opacity: 0 },
					'100%': { height: '100%', opacity: 1 },
				},
				grow: {
					//width grows from 0 to 100%
					'0%': { width: '0%', opacity: 0 },
					'20%': { width: '80%' },
					'80%': { width: '110%' },
					'100%': { width: '100%', opacity: 1 },
				},
			},
		},
	},
	plugins: [],
}
