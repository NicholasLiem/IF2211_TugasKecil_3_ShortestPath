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