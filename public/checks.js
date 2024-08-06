function checkConfPass() {
    var password = document.getElementById("password");
    var confirm_password = document.getElementById("confpassword");
    if (password.value != confirm_password.value) {
        confpassword.setCustomValidity("Passwords Don't Match");
    } else {
        confpassword.setCustomValidity('');
    }
}

function minPasswordLength() {
    var password = document.getElementById("password");
    if (password.value.length < 8) {
        password.setCustomValidity("Password must be at least 8 characters long");
    } else {
        password.setCustomValidity('');
    }
    checkConfPass();
}

function checkquant() {
    var quant = document.getElementById("quantity");
    quantity = parseInt(quant.value);
    if (quantity < 1 || isNaN(quantity)) {
        quant.setCustomValidity("Quantity must be at least 1");
    }
    if (quantity > 500) {
        quant.setCustomValidity("Quantity must be less than 500");
    }
}

function increment() {
    checkquant();
    var quant = document.getElementById("quant");
    var tobeupdated = document.getElementById("quantity");
    if (parseInt(quant.value) + parseInt(tobeupdated.value) > 500) {
        alert("Quantity must be less than 500");
    } else {
        final = parseInt(quant.value) + parseInt(tobeupdated.value);
        document.getElementById("send").value = final;
        quant.value = final;
    }
}

function decrement() {
    checkquant();
    var quant = document.getElementById("quant");
    var tobeupdated = document.getElementById("quantity");
    if (parseInt(quant.value) - parseInt(tobeupdated.value) < 1) {
        alert("Quanity must be at least 1");
    } else {
        final = parseInt(quant.value) - parseInt(tobeupdated.value);
        document.getElementById("send").value = final;
        quant.value = final;
    }
}