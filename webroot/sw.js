self.addEventListener('activate', function(event) {
    console.log('Service Worker Activate...');
    console.log(event);

    return self.clients.claim();
});

self.addEventListener('fetch', function(event) {
    if(event.request.url.includes(".ts")) {
        const url = event.request.url.split("/").splice(3); // ["hls", "0", "stream-foo<unix>.ts"]
        self.activeHlsStream = url[1];
        self.loadedHlsSegment = url.join("/");
        console.log(self.loadedHlsSegment);
    }
});