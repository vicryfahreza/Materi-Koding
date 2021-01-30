// macam macam popbox
// alert,prompt,confirm

// alert hanya menunjukan sebuah pesan 


// prompt untuk menuliskan kata untuk
// user dan mengembalikan nilai


// confirm sama seperti alert tapi membalikan 
// nilai berupa boolean

alert('selamat datang');
var lagi = true;

while(lagi){
    var nama = prompt("masukkan nama anda = ");
    alert('halo ' + nama);

    lagi = confirm("coba lagi?");

}


