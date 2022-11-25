export namespace internal {
	
	export class Message {
	    type: string;
	    content: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Message(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.content = source["content"];
	        this.name = source["name"];
	    }
	}

}

