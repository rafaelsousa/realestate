import React from 'react';
import { render, screen } from '@testing-library/react';
import PropertyForm from './App';

test('renders learn react link', () => {
  render(<PropertyForm />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
