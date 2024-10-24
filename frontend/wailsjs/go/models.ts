export namespace main {
	
	export class Cancel {
	    enabled: boolean;
	    entity: string;
	    op: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new Cancel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	    }
	}
	export class Config {
	    serverPort: string;
	    certPath: string;
	    keyPath: string;
	    databasePath: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.serverPort = source["serverPort"];
	        this.certPath = source["certPath"];
	        this.keyPath = source["keyPath"];
	        this.databasePath = source["databasePath"];
	    }
	}
	export class Delay {
	    enabled: boolean;
	    entity: string;
	    op: string;
	    value: string;
	    delaySec: number;
	
	    static createFrom(source: any = {}) {
	        return new Delay(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.delaySec = source["delaySec"];
	    }
	}
	export class Header {
	    action: string;
	    isRequest: boolean;
	    name: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new Header(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.action = source["action"];
	        this.isRequest = source["isRequest"];
	        this.name = source["name"];
	        this.value = source["value"];
	    }
	}
	export class HttpRequestLog {
	    // Go type: time
	    timestamp: any;
	    scheme: string;
	    method: string;
	    host: string;
	    path: string;
	    requestHeaders: {[key: string]: string[]};
	    responseHeaders: {[key: string]: string[]};
	    requestBody: string;
	    responseBody: string;
	    cancelled: boolean;
	    redirected: boolean;
	    requestHeaderModified: boolean;
	    responseHeaderModified: boolean;
	    requestBodyModified: boolean;
	    responseBodyModified: boolean;
	    delayed: number;
	
	    static createFrom(source: any = {}) {
	        return new HttpRequestLog(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.timestamp = this.convertValues(source["timestamp"], null);
	        this.scheme = source["scheme"];
	        this.method = source["method"];
	        this.host = source["host"];
	        this.path = source["path"];
	        this.requestHeaders = source["requestHeaders"];
	        this.responseHeaders = source["responseHeaders"];
	        this.requestBody = source["requestBody"];
	        this.responseBody = source["responseBody"];
	        this.cancelled = source["cancelled"];
	        this.redirected = source["redirected"];
	        this.requestHeaderModified = source["requestHeaderModified"];
	        this.responseHeaderModified = source["responseHeaderModified"];
	        this.requestBodyModified = source["requestBodyModified"];
	        this.responseBodyModified = source["responseBodyModified"];
	        this.delayed = source["delayed"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ModifyResponseBody {
	    enabled: boolean;
	    entity: string;
	    op: string;
	    value: string;
	    body: string;
	
	    static createFrom(source: any = {}) {
	        return new ModifyResponseBody(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.body = source["body"];
	    }
	}
	export class ModifyRequestBody {
	    enabled: boolean;
	    entity: string;
	    op: string;
	    value: string;
	    body: string;
	
	    static createFrom(source: any = {}) {
	        return new ModifyRequestBody(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.body = source["body"];
	    }
	}
	export class ModifyHeader {
	    enabled: boolean;
	    entity: string;
	    op: string;
	    value: string;
	    mods: Header[];
	
	    static createFrom(source: any = {}) {
	        return new ModifyHeader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.mods = this.convertValues(source["mods"], Header);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Redirect {
	    enabled: boolean;
	    entity: string;
	    op: string;
	    value: string;
	    toValue: string;
	
	    static createFrom(source: any = {}) {
	        return new Redirect(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enabled = source["enabled"];
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.toValue = source["toValue"];
	    }
	}
	export class InValue {
	    redirect: Redirect;
	    cancel: Cancel;
	    delay: Delay;
	    modifyHeader: ModifyHeader;
	    modifyRequestBody: ModifyRequestBody;
	    modifyResponseBody: ModifyResponseBody;
	
	    static createFrom(source: any = {}) {
	        return new InValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.redirect = this.convertValues(source["redirect"], Redirect);
	        this.cancel = this.convertValues(source["cancel"], Cancel);
	        this.delay = this.convertValues(source["delay"], Delay);
	        this.modifyHeader = this.convertValues(source["modifyHeader"], ModifyHeader);
	        this.modifyRequestBody = this.convertValues(source["modifyRequestBody"], ModifyRequestBody);
	        this.modifyResponseBody = this.convertValues(source["modifyResponseBody"], ModifyResponseBody);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	export class ReturnValue {
	    redirects: Redirect[];
	    cancels: Cancel[];
	    delays: Delay[];
	    modifyHeaders: ModifyHeader[];
	    modifyRequestBody: ModifyRequestBody[];
	    modifyResponseBody: ModifyResponseBody[];
	    logs: string[];
	    httpRequests: HttpRequestLog[];
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new ReturnValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.redirects = this.convertValues(source["redirects"], Redirect);
	        this.cancels = this.convertValues(source["cancels"], Cancel);
	        this.delays = this.convertValues(source["delays"], Delay);
	        this.modifyHeaders = this.convertValues(source["modifyHeaders"], ModifyHeader);
	        this.modifyRequestBody = this.convertValues(source["modifyRequestBody"], ModifyRequestBody);
	        this.modifyResponseBody = this.convertValues(source["modifyResponseBody"], ModifyResponseBody);
	        this.logs = source["logs"];
	        this.httpRequests = this.convertValues(source["httpRequests"], HttpRequestLog);
	        this.error = source["error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

