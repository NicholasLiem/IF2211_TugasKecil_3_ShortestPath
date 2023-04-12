const visualizer = document.getElementById("visualizer"),
    ctx = visualizer.getContext("2d"),
    nodeRadius = 7,
    coords = [],
    srcBtn = document.getElementById("src-btn"),
    dstBtn = document.getElementById("dst-btn"),
    srcName = document.getElementById("src-name"),
    dstName = document.getElementById("dst-name"),
    costLabel = document.getElementById("result-cost");

let graph,
    onHover = -1,
    inChoose = "",
    src = -1,
    dst = -1,
    route = [];

class Node {
    constructor({index, name, latitude, longitude}) {
        this.index = index;
        this.name = name;
        this.latitude = latitude;
        this.longitude = longitude;
    }
}

function exists(arrayNx2, element) {
    for (const [a, b] of arrayNx2) {
        if (a === element[0] && b === element[1]) {
            return true;
        }
    }
    return false;
}

class Graph {
    constructor({nodes, edges}) {
        this.nodesOrigin = nodes;
        this.nodes = [];
        for (let i = 0; i < Object.keys(nodes).length; i++) {
            const item = nodes[i];
            this.nodes[item["Index"]] = new Node({
                index: item["Index"],
                name: item["Name"],
                latitude: item["Latitude"],
                longitude: item["Longitude"],
            });
        }
        this.maxLat = this.nodes.reduce((prev, curr) => prev.latitude > curr.latitude ? prev : curr).latitude;
        this.minLat = this.nodes.reduce((prev, curr) => prev.latitude < curr.latitude ? prev : curr).latitude;
        this.maxLon = this.nodes.reduce((prev, curr) => prev.longitude > curr.longitude ? prev : curr).longitude;
        this.minLon = this.nodes.reduce((prev, curr) => prev.longitude < curr.longitude ? prev : curr).longitude;
        this.edges = edges;
    }

    draw(canvas) {
        if (!graph) {
            return
        }
        const ctx = canvas.getContext("2d");
        const width = canvas.width * 0.9;
        const height = canvas.height * 0.9;
        for (const node of this.nodes) {
            const x = canvas.width * 0.05 + (node.longitude - this.minLon) / (this.maxLon - this.minLon) * width;
            const y = canvas.height * 0.05 + (node.latitude - this.minLat) / (this.maxLat - this.minLat) * height;
            coords[node.index] = [x, y];
        }
        const drawn = [];
        for (const [from, edges] of Object.entries(this.edges)) {
            for (const [to, weight] of Object.entries(edges)) {
                if (!exists(drawn, [from, to])) {
                    const [x1, y1] = coords[from];
                    const [x2, y2] = coords[to];
                    const color = exists(route, [from, to]) ? "#ff8f00" : "#a0aaef";
                    drawLine(x1, y1, x2, y2, color, ctx);
                    drawText(weight.toString(), (x1 + x2) / 2 - 10, (y1 + y2) / 2 - 10, ctx);
                    drawn.push([from, to]);
                    drawn.push([to, from]);
                }
            }
        }
        for (const [idx, [x, y]] of coords.entries()) {
            const radius = nodeRadius + (onHover === idx ? 3 : 0);
            drawText(this.nodes[idx].name, x, y - radius - 10, ctx);
            const color = src === idx || dst === idx ? "#0062cc" : "#fff"
            drawCircle(x, y, radius, color, ctx);
        }
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

function drawCircle(x, y, radius, color, ctx) {
    ctx.beginPath();
    ctx.fillStyle = color;
    ctx.arc(x, y, radius, 0, 2 * Math.PI);
    ctx.fill();
}

function drawText(text, x, y, ctx) {
    ctx.font = "15px Poppins";
    ctx.textAlign = "center";
    ctx.fillStyle = "#fff"
    ctx.fillText(text, x, y)
}

function drawLine(x1, y1, x2, y2, color, ctx) {
    ctx.beginPath();
    ctx.strokeStyle = color;
    ctx.moveTo(x1, y1);
    ctx.lineTo(x2, y2);
    ctx.stroke();
}

visualizer.style.width = "100%";
visualizer.style.height = "75%";
visualizer.width = visualizer.offsetWidth;
visualizer.height = visualizer.offsetHeight;

visualizer.addEventListener("mousemove", ev => {
    const pos = getMousePos(visualizer, ev)
    if (graph) {
        for (const [idx, [x, y]] of coords.entries()) {
            const radius = nodeRadius + (onHover === idx ? 5 : 0);
            if (pos.x >= x - radius && pos.x <= x + radius && pos.y >= y - radius && pos.y <= y + radius) {
                visualizer.style.cursor = "pointer";
                if (onHover !== idx) {
                    onHover = idx;
                    ctx.clearRect(0, 0, visualizer.width, visualizer.height);
                    graph.draw(visualizer);
                }
                break;
            } else {
                if (onHover !== -1) {
                    onHover = -1;
                    ctx.clearRect(0, 0, visualizer.width, visualizer.height);
                    graph.draw(visualizer);
                }
                visualizer.style.cursor = "default";
            }
        }
    }
});

visualizer.addEventListener("mousedown", ev => {
    const pos = getMousePos(visualizer, ev)
    if (graph) {
        for (const [idx, [x, y]] of coords.entries()) {
            const radius = nodeRadius + (onHover === idx ? 5 : 0);
            if (pos.x >= x - radius && pos.x <= x + radius && pos.y >= y - radius && pos.y <= y + radius) {
                if (inChoose) {
                    if (inChoose === "src") {
                        src = idx;
                        srcBtn.style.filter = "brightness(100%)";
                        srcBtn.style.transform = "scale(100%)";
                        srcBtn.style.border = "0 none transparent";
                        srcName.innerText = graph.nodes[idx].name;
                    } else {
                        dst = idx;
                        dstBtn.style.filter = "brightness(100%)";
                        dstBtn.style.transform = "scale(100%)";
                        dstBtn.style.border = "0 none transparent";
                        dstName.innerText = graph.nodes[idx].name;
                    }
                    ctx.clearRect(0, 0, visualizer.width, visualizer.height);
                    graph.draw(visualizer);
                    inChoose = "";
                }
                break;
            }
        }
    }
});

window.addEventListener("resize", _ => {
    const width = visualizer.clientWidth;
    const height = visualizer.clientHeight;
    visualizer.width = width;
    visualizer.height = height;
    graph.draw(visualizer)
});

function displayError(message) {
    const error = document.getElementById("error")
    const msg = error.getElementsByClassName("error-msg")[0]
    msg.innerHTML = message
    error.style.visibility = "visible"
}

function hideError() {
    const error = document.getElementById("error")
    error.style.visibility = "hidden"
}

function displayControls() {
    const controls = document.getElementById("controls");
    controls.style.visibility = "visible";
}

function parseFile(file) {
    const reader = new FileReader();
    reader.onload = async (e) => {
        await fetch("http://localhost:8080/parse", {
            method: "POST",
            body: e.target.result
        }).then(async (response) => {
            if (response.ok) {
                if (response.status === 500) {
                    console.log(response);
                    throw new Error("Failed to parse file: " + await response.text());
                }
                return response.json();
            }
            throw new Error("Failed to upload file: " + response.statusText);
        }).then(async (responseJson) => {
            hideError();
            const {Nodes, Edges} = await responseJson;
            graph = new Graph({
                nodes: Nodes,
                edges: Edges,
            });
            dst = -1;
            src = -1;
            route = [];
            costLabel.style.visibility = "hidden";
            
            ctx.clearRect(0, 0, visualizer.width, visualizer.height);
            graph.draw(visualizer);
            displayControls();
        }).catch((error) => {
            displayError(error.message);
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

srcBtn.addEventListener("mousedown", (_) => {
    srcBtn.style.transform = "scale(95%)";
    srcBtn.style.filter = "brightness(95%)"
    srcBtn.style.border = "1px solid white";
    inChoose = "src";
    dstBtn.style.transform = "scale(100%)";
    dstBtn.style.filter = "brightness(100%)";
    dstBtn.style.border = "0 none transparent";
});

dstBtn.addEventListener("mousedown", (_) => {
    dstBtn.style.transform = "scale(95%)";
    dstBtn.style.filter = "brightness(95%)"
    dstBtn.style.border = "1px solid white";
    inChoose = "dst";
    srcBtn.style.transform = "scale(100%)";
    srcBtn.style.filter = "brightness(100%)";
    srcBtn.style.border = "0 none transparent";
});

document.getElementById("calculate-button").addEventListener("mousedown", async (_) => {
    srcBtn.style.filter = "brightness(100%)";
    srcBtn.style.transform = "scale(100%)";
    srcBtn.style.border = "0 none transparent";
    dstBtn.style.filter = "brightness(100%)";
    dstBtn.style.transform = "scale(100%)";
    dstBtn.style.border = "0 none transparent";
    if (src === -1) {
        displayError("No source selected")
        return
    }
    if (dst === -1) {
        displayError("No destination selected")
        return
    }
    const method = document.getElementById("algorithm").value;
    if (!method) {
        displayError("No algorithm selected")
        return
    }
    const body = {
        graph: {
            Nodes: graph.nodesOrigin,
            Edges: graph.edges
        },
        src,
        dst,
        method
    }
    await fetch("http://localhost:8080/search", {
        method: "POST", body: JSON.stringify(body)
    }).then(async (response) => {
        if (response.ok) {
            return response.json();
        }
        throw new Error("Failed to search path: " + await response.text());
    }).then(async (responseJson) => {
        route = [];
        const {Route, Cost} = await responseJson;
        for (let i = 0; i < Route.length - 1; i++) {
            const from = Route[i].toString();
            const to = Route[i + 1].toString();
            route.push(...[[from, to], [to, from]])
        }
        ctx.clearRect(0, 0, visualizer.width, visualizer.height);
        graph.draw(visualizer);
        costLabel.style.visibility = "visible";
        costLabel.innerText = "Total cost: " + Cost;
    }).catch((error) => {
        displayError(error.message);
    });
});
