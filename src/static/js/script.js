let visualizer = document.getElementById("visualizer");
let ctx = visualizer.getContext("2d");

class Node {
    constructor(index, name, latitude, longitude) {
        this.index = index;
        this.name = name;
        this.latitude = latitude;
        this.longitude = longitude;
    }
}

class Graph {
    constructor({nodes, edges}) {
        this.nodes = nodes;
        this.edges = edges;
    }
}

function getMousePos(canvas, evt) {
    const rect = canvas.getBoundingClientRect(), // abs. size of element
        scaleX = canvas.width / rect.width,    // relationship bitmap vs. element for x
        scaleY = canvas.height / rect.height;  // relationship bitmap vs. element for y

    return {
        x: (evt.clientX - rect.left) * scaleX,   // scale mouse coordinates after they have
        y: (evt.clientY - rect.top) * scaleY     // been adjusted to be relative to element
    };
}

function drawCircle(x, y, radius) {
    ctx.beginPath();
    ctx.arc(x, y, radius, 0, 2 * Math.PI);
    ctx.stroke();
}

visualizer.style.width = '100%';
visualizer.style.height = '100%';
visualizer.width = visualizer.offsetWidth;
visualizer.height = visualizer.offsetHeight;

// visualizer.addEventListener("mousemove", ev => {
//     const pos = getMousePos(visualizer, ev)
//     ctx.clearRect(0, 0, visualizer.width, visualizer.height)
//     drawCircle(pos.x, pos.y, 20)
// });

window.addEventListener("resize", ev => {
    const width = visualizer.clientWidth;
    const height = visualizer.clientHeight;
    visualizer.width = width;
    visualizer.height = height;
});

function displayError(message) {
    const error = document.getElementById("error")
    const msg = error.getElementsByClassName("error-msg")[0]
    msg.innerHTML = message
    error.style.visibility = "visible"
}

function hideError() {
    const error = document.getElementById("error")
    const canvas = document.getElementById("visualizer")
    canvas.style.backgroundColor = "#eee"
    error.style.visibility = "hidden"
}

function parseFile(file) {
    const reader = new FileReader();
    reader.onload = (e) => {
        fetch("http://localhost:8080/parse", {
            method: "POST",
            body: e.target.result
        }).then((response) => {
            if (response.ok) {
                if (response.status === 204) {
                    throw new Error("Failed to parse file")
                }
                return response.json()
            }
            throw new Error("Failed to upload file")
        }).then(async (responseJson) => {
            hideError()
            const {Nodes, Edges} = await responseJson
            console.log(new Graph({
                nodes: Nodes,
                edges: Edges,
            }))
        }).catch((error) => {
            displayError(error.message)
        })
    }
    reader.readAsDataURL(file);
    return false;
}

window.onload = function () {
    Particles.init({
        selector: ".background",
        color: "#a0aaef",
        connectParticles: true,
    });
};

function submitForm() {
    var form = document.querySelector('form');
    form.submit();
}

const urlParams = new URLSearchParams(window.location.search);
const uploaded = urlParams.get('uploaded');
if (uploaded === 'true') {
    alert('File uploaded successfully');
} else if (uploaded === 'false') {
    alert('File type is not supported, upload file failed')
}