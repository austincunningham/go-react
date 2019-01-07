import React, { Component } from 'react';
import Header from './components/header.js';
//import {c_background_image_BackgroundImage, c_background_image_BackgroundImage_xs} from '@patternfly/react-tokens';
//import { PageHeader } from '@patternfly/react-core';

//c_background_image_BackgroundImage.name === '--pf-c-background-image--BackgroundImage';
//c_background_image_BackgroundImage.value === 'url(assets/images/pfbg_576.jpg)';
//c_background_image_BackgroundImage.var === 'var(--pf-c-background-image--BackgroundImage)';

class App extends Component {
  render() {
    return (
      <div className="App">
         {Header}
      </div>
    );
  }
}

export default App;
