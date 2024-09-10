export namespace main {
	
	export class Redirect {
	    entity: string;
	    op: string;
	    value: string;
	    toType: string;
	    toValue: string;
	
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
	    }
	}
	export class ReturnValue {
	    redirects: Redirect[];
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new ReturnValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.redirects = this.convertValues(source["redirects"], Redirect);
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

