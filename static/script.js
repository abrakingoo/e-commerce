let images = document.querySelectorAll(".caro_img");
let indx = 0;

changeImg = () => {
    images.forEach(image => image.classList.remove("active"));

    images[indx].classList.add("active");

    indx = (indx + 1) % images.length;

    setTimeout(changeImg, 6000);
}

changeImg()