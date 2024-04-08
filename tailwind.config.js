module.exports = {
  content: ["./templates/*.{html,js,tmpl}"],
  theme: {
    extend: {
      animation: {
        fadeIn: "fadeIn .1s ease-in-out",
        fadeOut: "fadeOut .1s ease-in-out",
      },

      keyframes: {
        fadeIn: { from: { opacity: 0 }, to: { opacity: 1 } },
        fadeOut: { from: { opacity: 100 }, to: { opacity: 0 } },
      },
    },
  },
  plugins: [],
};
