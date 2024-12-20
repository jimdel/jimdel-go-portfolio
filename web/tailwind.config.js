/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ", "./views/**/*._templ"],
  theme: {
    extend: {
      backgroundImage: {
        hero: "url('/static/code.jpg')", // Custom background image
      },
    },
  },
  plugins: [],
};
