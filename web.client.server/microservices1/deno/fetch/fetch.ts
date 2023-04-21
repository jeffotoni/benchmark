const res = await fetch("http://localhost:3000/v1/avatar");
const body = await res.text();
console.log(body);