import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ReadOnlyTodoItemComponent } from './read-only-todo-item.component';

describe('ReadOnlyTodoItemComponent', () => {
  let component: ReadOnlyTodoItemComponent;
  let fixture: ComponentFixture<ReadOnlyTodoItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ReadOnlyTodoItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ReadOnlyTodoItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
