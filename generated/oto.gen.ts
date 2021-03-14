// Code generated by oto; DO NOT EDIT.

// HeadersFunc allows you to mutate headers for each request.
// Useful for adding authorization into the client.
interface HeadersFunc {
	(headers: HeadersInit);
}

// Client provides access to remote services.
export class Client {
	// basepath is the path prefix for the requests.
	// This may be a path, or an absolute URL.
	public basepath: String = '/oto/'
	// headers allows calling code to mutate the HTTP
	// headers of the underlying HTTP requests.
	public headers?: HeadersFunc
}


// GreeterService makes nice greetings.
export class GreeterService {
	constructor(readonly client: Client) {}
	
	// Greet makes a greeting.
	async greet(greetRequest: GreetRequest = null) {
		if (greetRequest == null) {
			greetRequest = new GreetRequest();
		}
		const headers: HeadersInit = new Headers();
		headers.set('Accept', 'application/json');
		headers.set('Content-Type', 'application/json');
		if (this.client.headers) {
			await this.client.headers(headers);
		}
		const response = await fetch(this.client.basepath + 'GreeterService.Greet', {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(greetRequest),
		})
		if (response.status !== 200) {
			throw new Error(`GreeterService.Greet: ${response.status} ${response.statusText}`);
		}
		return response.json().then((json) => {
			if (json.error) {
				throw new Error(json.error);
			}
			return new GreetResponse(json);
		})
	}
	
}

// TodoService create, read, update or delete.
export class TodoService {
	constructor(readonly client: Client) {}
	
		async create(createTodoRequest: CreateTodoRequest = null) {
		if (createTodoRequest == null) {
			createTodoRequest = new CreateTodoRequest();
		}
		const headers: HeadersInit = new Headers();
		headers.set('Accept', 'application/json');
		headers.set('Content-Type', 'application/json');
		if (this.client.headers) {
			await this.client.headers(headers);
		}
		const response = await fetch(this.client.basepath + 'TodoService.Create', {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(createTodoRequest),
		})
		if (response.status !== 200) {
			throw new Error(`TodoService.Create: ${response.status} ${response.statusText}`);
		}
		return response.json().then((json) => {
			if (json.error) {
				throw new Error(json.error);
			}
			return new CreateTodoResponse(json);
		})
	}
	
		async delete(deleteTodoRequest: DeleteTodoRequest = null) {
		if (deleteTodoRequest == null) {
			deleteTodoRequest = new DeleteTodoRequest();
		}
		const headers: HeadersInit = new Headers();
		headers.set('Accept', 'application/json');
		headers.set('Content-Type', 'application/json');
		if (this.client.headers) {
			await this.client.headers(headers);
		}
		const response = await fetch(this.client.basepath + 'TodoService.Delete', {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(deleteTodoRequest),
		})
		if (response.status !== 200) {
			throw new Error(`TodoService.Delete: ${response.status} ${response.statusText}`);
		}
		return response.json().then((json) => {
			if (json.error) {
				throw new Error(json.error);
			}
			return new DeleteTodoResponse(json);
		})
	}
	
		async get(getTodoRequest: GetTodoRequest = null) {
		if (getTodoRequest == null) {
			getTodoRequest = new GetTodoRequest();
		}
		const headers: HeadersInit = new Headers();
		headers.set('Accept', 'application/json');
		headers.set('Content-Type', 'application/json');
		if (this.client.headers) {
			await this.client.headers(headers);
		}
		const response = await fetch(this.client.basepath + 'TodoService.Get', {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(getTodoRequest),
		})
		if (response.status !== 200) {
			throw new Error(`TodoService.Get: ${response.status} ${response.statusText}`);
		}
		return response.json().then((json) => {
			if (json.error) {
				throw new Error(json.error);
			}
			return new GetTodoResponse(json);
		})
	}
	
		async getAll(getAllTodosRequest: GetAllTodosRequest = null) {
		if (getAllTodosRequest == null) {
			getAllTodosRequest = new GetAllTodosRequest();
		}
		const headers: HeadersInit = new Headers();
		headers.set('Accept', 'application/json');
		headers.set('Content-Type', 'application/json');
		if (this.client.headers) {
			await this.client.headers(headers);
		}
		const response = await fetch(this.client.basepath + 'TodoService.GetAll', {
			method: 'POST',
			headers: headers,
			body: JSON.stringify(getAllTodosRequest),
		})
		if (response.status !== 200) {
			throw new Error(`TodoService.GetAll: ${response.status} ${response.statusText}`);
		}
		return response.json().then((json) => {
			if (json.error) {
				throw new Error(json.error);
			}
			return new GetAllTodosResponse(json);
		})
	}
	
}



export class CreateTodoRequest {
	constructor(data?: any) {
		if (data) {
		
			
			this.title = data.title;
			
		
			
			this.description = data.description;
			
		
		}
	}

		title: string;

		description: string;

}

export class Todo {
	constructor(data?: any) {
		if (data) {
		
			
			this.id = data.id;
			
		
			
			this.title = data.title;
			
		
			
			this.description = data.description;
			
		
			
			this.completed = data.completed;
			
		
			
			this.createdAt = data.createdAt;
			
		
			
			this.updatedAt = data.updatedAt;
			
		
		}
	}

		id: number;

		title: string;

		description: string;

		completed: boolean;

		createdAt: string;

		updatedAt: string;

}

export class CreateTodoResponse {
	constructor(data?: any) {
		if (data) {
		
			
				
					this.todo = new Todo(data.todo);
				
			
		
			
			this.error = data.error;
			
		
		}
	}

		todo: todo.Todo;

	// Error is string explaining what went wrong. Empty if everything was fine.
	error: string;

}

export class DeleteTodoRequest {
	constructor(data?: any) {
		if (data) {
		
			
			this.todoID = data.todoID;
			
		
		}
	}

		todoID: number;

}

export class DeleteTodoResponse {
	constructor(data?: any) {
		if (data) {
		
			
			this.success = data.success;
			
		
			
			this.error = data.error;
			
		
		}
	}

		success: boolean;

	// Error is string explaining what went wrong. Empty if everything was fine.
	error: string;

}

// GetAllTodosRequest - needs pagination
export class GetAllTodosRequest {
	constructor(data?: any) {
		if (data) {
		
		}
	}

}

// GetAllTodosResponse - needs pagination
export class GetAllTodosResponse {
	constructor(data?: any) {
		if (data) {
		
			
				
					if (data.todos) {
						this.todos = new Array<Todo>()
						for (let i = 0; i < data.todos.length; i++) {
							this.todos.push(new Todo(data.todos[i]));
						}
					}
				
			
		
			
			this.error = data.error;
			
		
		}
	}

		todos: todo.Todo[];

	// Error is string explaining what went wrong. Empty if everything was fine.
	error: string;

}

// GetTodoRequest based by id
export class GetTodoRequest {
	constructor(data?: any) {
		if (data) {
		
			
			this.id = data.id;
			
		
		}
	}

		id: number;

}

// GetTodoResponse returns the todo.
export class GetTodoResponse {
	constructor(data?: any) {
		if (data) {
		
			
				
					this.todo = new Todo(data.todo);
				
			
		
			
			this.error = data.error;
			
		
		}
	}

		todo: todo.Todo;

	// Error is string explaining what went wrong. Empty if everything was fine.
	error: string;

}

// GreetRequest is the request object for GreeterService.Greet.
export class GreetRequest {
	constructor(data?: any) {
		if (data) {
		
			
			this.name = data.name;
			
		
		}
	}

	// Name is the person to greet.
	name: string;

}

// GreetResponse is the response object containing a person's greeting.
export class GreetResponse {
	constructor(data?: any) {
		if (data) {
		
			
			this.greeting = data.greeting;
			
		
			
			this.error = data.error;
			
		
		}
	}

	// Greeting is the greeting that was generated.
	greeting: string;

	// Error is string explaining what went wrong. Empty if everything was fine.
	error: string;

}

