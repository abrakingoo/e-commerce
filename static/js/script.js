// carrosel loading
document.addEventListener('DOMContentLoaded', () => {
    let images = document.querySelectorAll(".caro_img");
    if (images.length === 0) {
        console.warn("No carousel images found!");
        return;
    }

    let indx = 0;

    const changeImg = () => {
        images.forEach(image => image.classList.remove("active"));
        images[indx].classList.add("active");
        indx = (indx + 1) % images.length; // Loop back to 0 after the last image
        setTimeout(changeImg, 6000);
    }
    changeImg();
});


// 
// 
let account = document.querySelector(".account");
let actions = document.getElementById("actions");
let body = document.body;

account.addEventListener("click", (event) => {
    event.stopPropagation();
    if (getComputedStyle(actions).display === "none") {
        actions.style.display = "block";
    } else {
        actions.style.display = "none";
    }
});

body.addEventListener("click", () => {
    if (getComputedStyle(actions).display === "block") {
        actions.style.display = "none";
    }
});


//
// 
// Listens And adds items to cart
let btns = document.querySelectorAll("#add_to_cart");
let count = document.querySelector(".items_count");

btns.forEach((btn) => {
    btn.addEventListener('click', (e) => {
        let id = btn.nextElementSibling.value;
        let cartItems = JSON.parse(localStorage.getItem('cartItems')) || [];

        if (btn.textContent == "Add To Cart") {
            // Add the item to the cart and update the localStorage
            fetch(`/cart?id=${id}`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                }
            })
            .then(res => {
                if (res.status === 401) {
                    count.innerHTML = 0;
                    alert("Please login to continue")
                } else if (res.ok) {
                    res.json().then(data => {
                        count.innerHTML = data.Count;
                        cartItems.push(id); // Add item ID to cart items array
                        localStorage.setItem('cartItems', JSON.stringify(cartItems)); // Store in localStorage
                    });
                    btn.textContent = "Remove"; // Update button text
                } else {
                    console.log(`Unexpected status code: ${res.status}`);
                }
            })
            .catch(err => {
                console.error("Error:", err);  // Log any network or fetch errors
            });
        } else if (btn.textContent == "Remove") {
            // Remove the item from the cart and update the localStorage
            fetch(`/remove?id=${id}`, {
                method: "POST",
                headers: { "Content-Type" : "application/json" },
            })
            .then(res => {
                if (res.ok) {
                    btn.textContent = "Add To Cart"; // Reset button text
                    cartItems = cartItems.filter(item => item !== id); // Remove item ID from array
                    localStorage.setItem('cartItems', JSON.stringify(cartItems)); // Update localStorage
                    window.location.reload();
                }
            })
        }
    });
});


// On page load, update the button text based on cart state
document.addEventListener("DOMContentLoaded", () => {
    let btns = document.querySelectorAll("#add_to_cart");
    let cartItems = JSON.parse(localStorage.getItem('cartItems')) || [];

    btns.forEach((btn) => {
        let id = btn.nextElementSibling.value;

        if (cartItems.includes(id)) {
            btn.textContent = "Remove"; // Item is in cart, show "Remove"
        } else {
            btn.textContent = "Add To Cart"; // Item is not in cart, show "Add To Cart"
        }
    });
});


let removeBtns = document.querySelectorAll("#remove_from_cart");

removeBtns.forEach((btn) => {
    let cartItems = JSON.parse(localStorage.getItem('cartItems')) || [];
    btn.addEventListener('click', ()=>{
        let id = btn.nextElementSibling.value
        fetch(`/remove?id=${id}`, {
            method: "POST",
            headers: {"Content-Type" : "application/json"},
        })
        .then(res => {
            if(res.ok) {
                cartItems = cartItems.filter(item => item !== id); // Remove item ID from array
                localStorage.setItem('cartItems', JSON.stringify(cartItems)); // Update localStorage
                window.location.reload();
            } 
            })
    });
});

let logOutBtn = document.getElementById("logout");

logOutBtn.addEventListener('click', (e) => {
    e.stopPropagation(); // Prevents the event from bubbling up
    localStorage.removeItem("cartItems"); // Removes only the 'cartItems' entry from localStorage
});

let chekoutBtn = document.getElementById("buy_btn") 
if (chekoutBtn != null) {
    chekoutBtn.addEventListener('click', ()=> {
        localStorage.removeItem("cartItems");
    })
} 

let hide = document.querySelector(".hide");  // Targeting button inside .hide
let el = document.querySelector("#profile_div #left"); 
let right = document.querySelector("#profile_div #right");  // Select #left inside #profile_div

hide.addEventListener('click', () => {
    el.style.display = "none";  // Hide the left section
    right.style.width = "100%"
});