export namespace main {
	
	export class Config {
	    serverPort: string;
	    certPath: string;
	    keyPath: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.serverPort = source["serverPort"];
	        this.certPath = source["certPath"];
	        this.keyPath = source["keyPath"];
	    }
	}
	export class Delay {
	    entity: string;
	    op: string;
	    value: string;
	    toType: string;
	    delaySec: number;
	
	    static createFrom(source: any = {}) {
	        return new Delay(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.toType = source["toType"];
	        this.delaySec = source["delaySec"];
	    }
	}
	export class Cancel {
	    entity: string;
	    op: string;
	    value: string;
	    toType: string;
	
	    static createFrom(source: any = {}) {
	        return new Cancel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.toType = source["toType"];
	    }
	}
	export class Redirect {
	    entity: string;
	    op: string;
	    value: string;
	    toType: string;
	    toValue: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Redirect(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	        this.toType = source["toType"];
	        this.toValue = source["toValue"];
	        this.enabled = source["enabled"];
	    }
	}
	export class InValue {
	    redirect: Redirect;
	    // Go type: Cancel
	    cancel: any;
	    delay: Delay;
	
	    static createFrom(source: any = {}) {
	        return new InValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.redirect = this.convertValues(source["redirect"], Redirect);
	        this.cancel = this.convertValues(source["cancel"], null);
	        this.delay = this.convertValues(source["delay"], Delay);
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
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new ReturnValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.redirects = this.convertValues(source["redirects"], Redirect);
	        this.cancels = this.convertValues(source["cancels"], Cancel);
	        this.delays = this.convertValues(source["delays"], Delay);
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

