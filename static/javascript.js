var colorBase = "Black",
    Element = document.querySelector('span');

function changeColor(val) {
    var color = "Black";
    if (val === "Fire") {
        color = "Red";
    } else {
        color = "Black";
    }

    spanText.style.color = color;
}

changeColor(colorBase);