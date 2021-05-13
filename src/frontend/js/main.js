import React from 'react'

import ReactDOM from 'react-dom';

import {
  Container,
  Row
} from 'react-bootstrap';

import Header from './components/header.js';

import {
  Tab,
  CodeArea,
  Editor
} from './components/main_editor.js';

var myCSS = require('../styles/main.css')

var bootstrapCSS = require('bootstrap/dist/css/bootstrap.css')

const App = () => {
  return (
    <Container fluid>
      <Row>
        <Header />
      </Row>
      <Row id="viewport" >
        <Editor>
          <Tab><CodeArea /></Tab>  
        </Editor>
      </Row>
      <div id="floater" style={{display: "none"}}></div>
    </Container>
  );
};

ReactDOM.render(App(), document.getElementById("root"));
