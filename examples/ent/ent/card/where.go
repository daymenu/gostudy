// Code generated by entc, DO NOT EDIT.

package card

import (
	"github.com/daymenu/gostudy/examples/ent/ent/predicate"
	"github.com/daymenu/gostudy/examples/ent/ent/schema"
	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// CardID applies equality check predicate on the "cardID" field. It's identical to CardIDEQ.
func CardID(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardID), vc))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// CardIDEQ applies the EQ predicate on the "cardID" field.
func CardIDEQ(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardID), vc))
	})
}

// CardIDNEQ applies the NEQ predicate on the "cardID" field.
func CardIDNEQ(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCardID), vc))
	})
}

// CardIDIn applies the In predicate on the "cardID" field.
func CardIDIn(vs ...schema.CardID) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCardID), v...))
	})
}

// CardIDNotIn applies the NotIn predicate on the "cardID" field.
func CardIDNotIn(vs ...schema.CardID) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCardID), v...))
	})
}

// CardIDGT applies the GT predicate on the "cardID" field.
func CardIDGT(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCardID), vc))
	})
}

// CardIDGTE applies the GTE predicate on the "cardID" field.
func CardIDGTE(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCardID), vc))
	})
}

// CardIDLT applies the LT predicate on the "cardID" field.
func CardIDLT(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCardID), vc))
	})
}

// CardIDLTE applies the LTE predicate on the "cardID" field.
func CardIDLTE(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCardID), vc))
	})
}

// CardIDContains applies the Contains predicate on the "cardID" field.
func CardIDContains(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCardID), vc))
	})
}

// CardIDHasPrefix applies the HasPrefix predicate on the "cardID" field.
func CardIDHasPrefix(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCardID), vc))
	})
}

// CardIDHasSuffix applies the HasSuffix predicate on the "cardID" field.
func CardIDHasSuffix(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCardID), vc))
	})
}

// CardIDEqualFold applies the EqualFold predicate on the "cardID" field.
func CardIDEqualFold(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCardID), vc))
	})
}

// CardIDContainsFold applies the ContainsFold predicate on the "cardID" field.
func CardIDContainsFold(v schema.CardID) predicate.Card {
	vc := string(v)
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCardID), vc))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		p(s.Not())
	})
}
