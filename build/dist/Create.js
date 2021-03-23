import {useState} from "../_snowpack/pkg/preact/hooks.js";
import {client} from "./App.js";
import {TodoService} from "./types/oto.gen.js";
function Create({onSubmit}) {
  const [title, setTitle] = useState("");
  const handleSubmit = async () => {
    const service = new TodoService(client);
    const res = await service.create({
      title,
      description: "",
      completed: false
    });
    if (res.error !== "")
      alert(res.error);
    onSubmit();
  };
  return /* @__PURE__ */ React.createElement("div", null, /* @__PURE__ */ React.createElement("form", {
    onSubmit: handleSubmit
  }, /* @__PURE__ */ React.createElement("input", {
    type: "text",
    name: "title",
    onChange: (e) => setTitle(e.currentTarget.value)
  }), /* @__PURE__ */ React.createElement("button", {
    onClick: handleSubmit
  }, "Submit")));
}
export {Create};
