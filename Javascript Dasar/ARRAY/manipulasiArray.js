// // menambah array
// var arr = ['d1','d2','s1','s2'];
// arr[0] = 'berubah';
// console.log(arr[0]);

// // menghapus array
// var arr = [1,2,3,4,5,'oke'];
// arr[0] = undefined;
// console.log(arr)


// // memanggil array
// var arr = ['dody','bowo','prabow',true,false];
// for(var i = 0; i < arr.length; i++){
//     console.log('mahasiswa ke-' + (i+1) + ':' + arr[i]);
// }


// var arr = ['tino', 'toni' , 'joni' , 'wanto']
// // join
// console.log(arr.join(' - '));


// // push
// var puss = ['sandika','galih','dody'];
// puss.push('riski');
// console.log(puss.join(' - '));

// // pop
// var popp = ['dika','gerald','niessa','jeff'];
// popp.pop('riski');
// console.log(popp.join(' - '));

// shift

// unshift

// splice
// splice(index awal,mau berapa,'elemenbaru1',....)
// var arr = ['d1','d2','s1','s2'];
// arr.splice(3,0,'dono','dina');
// console.log(arr.join(' - '));

// slice
// var arr = ['d1','d2','s1','s2'];
// var arr2 = arr.slice(1,3);
// console.log(arr.join(' - '));
// console.log(arr2.join(' - '));

// foreach
// var arr = ['d1','d2','s1','s2'];
// arr.forEach(function(e,i){
//     console.log('gelar ke' + i + '=' + e);
// });


// map
// var arr = [1,2,3,4,5,6,7,8,9];
// var arr2 = arr.map(function(e){
//     return e * 2;
// });
// console.log(arr2.join(' - '));

// sort
var arr = [1,4,2,1,11,23,51,9,10,5,6];
arr.sort(function(a,b){
    return a - b;
});
console.log(arr.join(' - '));



