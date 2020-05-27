export function simpleFetch(url, method, data=null) {
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open(method, url, false);
    xmlHttp.send(data);
    return xmlHttp.responseText;
}