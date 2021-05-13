import React from 'react';

import {
  Col,
  Jumbotron
} from 'react-bootstrap';

const CodeArea = React.forwardRef((props, ref) => (
  <Jumbotron fluid suppressContentEditableWarning={true} contentEditable onInput={props.onInput} ref={ref}>asdfasdfafa</Jumbotron>
));

const Tab = (props) => {
  return (
    <Container fluid></Container>
  );
};

class Editor extends React.Component {

  render() {
    return (
        <CodeArea ref={this.codeArea} onInput={this.something.bind(this)} />
    );
  };

  constructor(props) {
    super(props);
    this.state = {
      contents: "",
      name: "",
      tree: {},
    };
    this.componentDidUpdate = this.componentDidUpdate.bind(this);
    this.codeArea = React.createRef();
  }

  pushContents(name, contents) {
    const body = {
      action: "parseYAML",
      data: {
        name: name,
        contents: contents,
      },
    };
    const data = {
      method: "POST",
      headers: {
        "Content-type": "application/json",
      },
      body: JSON.stringify(body),
    };
    fetch('/ajax', data)
      .then(res => res.json())
      .then(
        data => {
          this.setState({
            contents: contents,
            tree: data,
          });
        }
      );
  }

  componentDidUpdate() {
    this.codeArea.current.content = this.state.contents;
  }

  loadContents(name, contents) {
    this.setState({
      name: name,
      contents: contents,
    });
    this.pushContents(name, contents);
  }

  something() {
    console.log("something");
  }

}

export {
  Editor,
  CodeArea,
  Tab
};
