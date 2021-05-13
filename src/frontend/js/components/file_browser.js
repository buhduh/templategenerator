import React from 'react';
import ReactDOM from 'react-dom';

//props.browserRef is the ref to the input of type file, originally used so i could "click" it
//this.renderElementRef MUST implement the loadContents method that takes a filename and the
//file contents
class FileBrowser extends React.Component {
  constructor(props) {
    super(props);
    //probably the <input type=file> element
    this.input = props.browserRef;
    //the file browser is a child of container, responsible for cleaning itself up
    this.container = props.container;
    //if something needs rendering upon file selection, this is it
    this.renderElement = props.renderElement;
    //has form func(name, contents) where name is the name of the file and contents is its contents
    this.renderElementRef = props.renderElementRef;
  }

  cleanup() {
    ReactDOM.unmountComponentAtNode(this.container);
  }

  render() {
    return <
      input 
      ref={this.input}
      style={{display: 'none'}} 
      id="file_browser" 
      type="file" 
      onChange={
        () => { 
          var fileName = this.input.current.files[0].name
          console.log(`loaded ${fileName}`);
          if(!!this.renderElement) {
            //If i want to reuse this, ill probably need to parameterize the second argument
            ReactDOM.render(this.renderElement, document.getElementById("canvas"));
          }
          let file = this.input.current.files[0];
          let reader = new FileReader();
          reader.readAsText(file, "UTF-8");
          reader.onload = function (evt) {
            this.renderElementRef.current.loadContents(
              fileName, evt.target.result
            );
          }.bind(this);
          this.cleanup();
        }
      }
    />;
  };
}

export default FileBrowser;
/*
          globals.currDocument.update()
  update(file) {
    let reader = new FileReader();
    reader.readAsText(file, "UTF-8");
    reader.onload = function (evt) {
      const body = {
        action: "newDocument",
        data: {
          contents: evt.target.result,
          name: file.name,
        }
      }
      const data = {
        method: "POST",
        headers: {
          "Content-type": "application/json"
        },
        body: JSON.stringify(body)
      }
      fetch('/ajax', data)
        .then(res => res.json())
        .then(data => console.log(data));
    }
  }
};
  update(file) {
    let reader = new FileReader();
    reader.readAsText(file, "UTF-8");
    reader.onload = function (evt) {
      const body = {
        action: "newDocument",
        data: {
          contents: evt.target.result,
          name: file.name,
        }
      }
      const data = {
        method: "POST",
        headers: {
          "Content-type": "application/json"
        },
        body: JSON.stringify(body)
      }
      fetch('/ajax', data)
        .then(res => res.json())
        .then(data => console.log(data));
    }
  }
};
*/
