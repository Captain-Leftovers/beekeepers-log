/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./internal/view/**/*.templ'],
	theme: {
		extend: {
			colors: {
				'cover': '#050505',
			},
		},
	},
	plugins: [],
}
