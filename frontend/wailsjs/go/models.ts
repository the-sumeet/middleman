export namespace main {
	
	export class Cancel {
	    entity: string;
	    op: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new Cancel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.entity = source["entity"];
	        this.op = source["op"];
	        this.value = source["value"];
	    }
	}
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
	    delaySec: number;
	
	    static createFrom(source: any = {}) {
	        return new Delay(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
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
	export class ModifyHeader {
	    entity: string;
	    op: string;
	    value: string;
	    mods: Header[];
	
	    static createFrom(source: any = {}) {
	        return new ModifyHeader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
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
	    cancel: Cancel;
	    delay: Delay;
	    modifyHeader: ModifyHeader;
	
	    static createFrom(source: any = {}) {
	        return new InValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.redirect = this.convertValues(source["redirect"], Redirect);
	        this.cancel = this.convertValues(source["cancel"], Cancel);
	        this.delay = this.convertValues(source["delay"], Delay);
	        this.modifyHeader = this.convertValues(source["modifyHeader"], ModifyHeader);
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

