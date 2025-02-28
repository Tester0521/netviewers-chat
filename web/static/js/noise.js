
const canvas = document.getElementById('noiseCanvas');
const ctx = canvas.getContext('2d');

canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

let imageData = ctx.createImageData(canvas.width, canvas.height);
let data = imageData.data;

const createNoise = () => {
    for (let i = 0; i < data.length; i += 4) {
        const value = Math.random() * 255; 

        data[i] = value;     
        data[i + 1] = value; 
        data[i + 2] = value; 
        data[i + 3] = 255 
    }
}

setInterval(() => {
    createNoise()
    ctx.putImageData(imageData, 0, 0);
}, 100)

window.addEventListener('resize', () => {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    imageData = ctx.createImageData(canvas.width, canvas.height);
    data = imageData.data
});