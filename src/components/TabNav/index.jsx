import React from 'react';
import PropTypes from 'prop-types';
import { Tag } from '@trussworks/react-uswds';
import { Tab, Tabs, TabList } from 'react-tabs';
import classNames from 'classnames/bind';
import styles from './index.module.scss';

const cx = classNames.bind(styles);

const TabNav = ({ options, children }) => (
  <Tabs>
    <TabList className={cx('tab-nav')}>
      {options.map(({ title, active, notice }, index) => (
        <Tab key={index.toString()} className={cx('tab-item')}>
          <span className={cx('tab-title', { 'tab-title-active': active })}>{title}</span>
          {notice && <Tag>{notice}</Tag>}
        </Tab>
      ))}
    </TabList>
    {children}
  </Tabs>
);

TabNav.propTypes = {
  options: PropTypes.arrayOf(
    PropTypes.shape({
      title: PropTypes.string,
      active: PropTypes.bool,
      notice: PropTypes.string,
    }),
  ).isRequired,
  children: PropTypes.arrayOf(PropTypes.node).isRequired,
};

export default TabNav;
