export namespace main {
	
	export class Config {
	    proxyServerPort: string;
	    webServerPort: string;
	    certPath: string;
	    keyPath: string;
	    databasePath: string;
	    webServerPath: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.proxyServerPort = source["proxyServerPort"];
	        this.webServerPort = source["webServerPort"];
	        this.certPath = source["certPath"];
	        this.keyPath = source["keyPath"];
	        this.databasePath = source["databasePath"];
	        this.webServerPath = source["webServerPath"];
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
	export class Rule {
	    id: any;
	    type: string;
	    enabled: boolean;
	    entity: string;
	    op: string;
	    value: string;
	    redirectTo: string;
	    delaySec: number;
	    requestBody: string;
	    responseBody: string;
	    requestHeaderMods: Header[];
	    responseHeaderMods: Header[];
	
	    static createFrom(source: any = {}) {
	        return new Rule(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.type = source["type"];
	        this.enabled = source["enabled"];
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.redirectTo = source["redirectTo"];
	        this.delaySec = source["delaySec"];
	        this.requestBody = source["requestBody"];
	        this.responseBody = source["responseBody"];
	        this.requestHeaderMods = this.convertValues(source["requestHeaderMods"], Header);
	        this.responseHeaderMods = this.convertValues(source["responseHeaderMods"], Header);
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
	export class InValue {
	    id: any;
	    rule: Rule;
	
	    static createFrom(source: any = {}) {
	        return new InValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.rule = this.convertValues(source["rule"], Rule);
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
	    insertedId: any;
	    rules: Rule[];
	    logs: string[];
	    httpRequests: HttpRequestLog[];
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new ReturnValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.insertedId = source["insertedId"];
	        this.rules = this.convertValues(source["rules"], Rule);
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

