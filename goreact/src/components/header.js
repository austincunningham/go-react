import React from 'react';
//import { Page, PageHeader, PageSidebar, PageSection, PageSectionVariants } from '@patternfly/react-core';

export const Header = (
    < header>Text goes here</header>
  );

export default Header;
// class VerticalPage extends React.Component {
//   constructor(props) {
//     super(props);
//     this.state = {
//       isNavOpen: true
//     };
//   }

//   onNavToggle = () => {
//     this.setState({
//       isNavOpen: !this.state.isNavOpen
//     });
//   };

//   render() {
//     const { isNavOpen } = this.state;

//     const logoProps = {
//       href: 'https://patternfly.org',
//       onClick: () => console.log('clicked logo'),
//       target: '_blank'
//     };
//     const Header = (
//       <PageHeader
//         logo="Logo"
//         logoProps={logoProps}
//         toolbar="Toolbar"
//         avatar=" | Avatar"
//         showNavToggle
//         onNavToggle={this.onNavToggle}
//       />
//     );
//     const Sidebar = <PageSidebar nav="Navigation" isNavOpen={isNavOpen} />;

//     return (
//       <Page header={Header} sidebar={Sidebar}>
//         <PageSection variant={PageSectionVariants.darker}>Section with darker background</PageSection>
//         <PageSection variant={PageSectionVariants.dark}>Section with dark background</PageSection>
//         <PageSection variant={PageSectionVariants.light}>Section with light background</PageSection>
//       </Page>
//     );
//   }
// }

// export default VerticalPage;
