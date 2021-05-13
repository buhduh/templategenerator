import React from 'react';
import ReactDOM from 'react-dom';

import FileBrowser from "./file_browser.js";
import Editor from "./main_editor.js";

import {
  Nav,
  NavDropdown
} from 'react-bootstrap';

class Header extends React.Component {

  constructor(props) {
    super(props)
    this.fileSelect = this.fileSelect.bind(this);
    this.browserRef = React.createRef();;
  }

  //TODO, cannot, for the life of me, get the "active" class to go away.
  //There's some sort of bubbling/capturing going on that im just picking up on
  fileSelect(eventKey, e) {
    e.target.classList.remove("active");
    switch(eventKey) {
      case "new_file":
        console.log(eventKey);
        break;
      case "open_file":
        const element = document.getElementById("floater");
        const editorRef = React.createRef();
        const editor = <Editor ref={editorRef} fileName="" contents="" />;
        const fileBrowser = <
          FileBrowser container={element}
          renderElement={editor}
          renderElementRef={editorRef}
          browserRef={this.browserRef}
        />;
        ReactDOM.render(fileBrowser, element);
        this.browserRef.current.click();
        break;
      case "save":
        console.log(eventKey);
        break;
      case "default":
        console.log(eventKey);
        break;
      default:
        console.error(`unrecognized eventKey: ${eventKey}`);
        break;
    }
  }

  render() {
    return (
      <Nav>
        <Nav.Item>
          <NavDropdown onSelect={this.fileSelect} title="File">
            <NavDropdown.Item eventKey="new_file">New</NavDropdown.Item>
            <NavDropdown.Item eventKey="open_file">Open</NavDropdown.Item>
            <NavDropdown.Item eventKey="save">Save</NavDropdown.Item>
            <NavDropdown.Item eventKey="save_as">Save As</NavDropdown.Item>
          </NavDropdown>
        </Nav.Item>
        <NavDropdown title="AWS">
        </NavDropdown>
        <NavDropdown title="Tools">
        </NavDropdown>
      </Nav>
    );
  }
}

export default Header;
