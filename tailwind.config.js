module.exports = {
  content: ["./templates/*.{html,js,tmpl}"],
  theme: {
    extend: {
      animation: {
        fadeIn: "fadeIn .2s ease-in-out",
        fadeOut: "fadeOut .2s ease-in-out",
      },

      keyframes: {
        fadeIn: { from: { opacity: 0 }, to: { opacity: 1 } },
        fadeOut: { from: { opacity: 100 }, to: { opacity: 0 } },
      },
    },
  },
  plugins: [],
};
