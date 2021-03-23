import {useEffect, useState, useCallback} from "../_snowpack/pkg/preact/hooks.js";
import {Create} from "./Create.js";
import {Client, TodoService} from "./types/oto.gen.js";
export const client = new Client();
function App() {
  const [todos, setTodos] = useState([]);
  const fetchData = useCallback(async () => {
    const service = new TodoService(client);
    const tt = await service.getAll();
    setTodos(tt.todos ?? []);
  }, []);
  useEffect(() => {
    fetchData();
  }, []);
  const onSubmit = () => {
    fetchData();
  };
  return /* @__PURE__ */ React.createElement("div", {
    className: "App"
  }, /* @__PURE__ */ React.createElement(Create, {
    onSubmit
  }), !todos ? /* @__PURE__ */ React.createElement("p", null, "Could not fetch todos") : !todos.length ? /* @__PURE__ */ React.createElement("p", null, "No todos created") : todos.map((t) => /* @__PURE__ */ React.createElement("div", {
    key: t.id
  }, t.title)));
}
export default App;
