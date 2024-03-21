let menuButton = document.querySelector(".button-menu");
let layout = document.querySelector(".layout");
let pageContent = document.querySelector(".page-content");
let responsiveBreakpoint = 991;

if (window.innerWidth <= responsiveBreakpoint) {
  layout.classList.add("nav-closed");
}

menuButton.addEventListener("click", function () {
  layout.classList.toggle("nav-closed");
});

pageContent.addEventListener("click", function () {
  if (window.innerWidth <= responsiveBreakpoint) {
    layout.classList.add("nav-closed");
  }
});


window.addEventListener("resize", function () {
  if (window.innerWidth > responsiveBreakpoint) {
    layout.classList.remove("nav-closed");
  }
});
