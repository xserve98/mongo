package mongo

// Define embedded types that mimic mgo.Collection, mgo.Database and
// mgo.Query. They are embedded to use the interfaces defined on this
// package, allowing mocking and correct testing of them.

import (
	"gopkg.in/mgo.v2"
	"time"
)

// ChangeInfo it's an embedded type of mgo.ChangeInfo, made to reduce
// importing to this package only, since it's necessary on other
// functions.
type ChangeInfo struct {
	*mgo.ChangeInfo
}

// NewChangeInfo creates a new embedded ChangeInfo, with all attributes
// for class.
func NewChangeInfo(u, r, m int, id interface{}) (c *ChangeInfo) {
	c = &ChangeInfo{
		ChangeInfo: &mgo.ChangeInfo{
			Updated:    u,
			Removed:    r,
			Matched:    m,
			UpsertedId: id,
		},
	}
	return
}

// NewRemoveInfo creates a new embedded ChangeInfo, with all attributes
// returned on calls to Remove.
func NewRemoveInfo(r int) (c *ChangeInfo) {
	c = NewChangeInfo(0, r, 0, nil)
	return
}

// NewUpdateInfo creates a new embedded ChangeInfo, with all attributes
// returned on calls to Update.
func NewUpdateInfo(u, m int) (c *ChangeInfo) {
	c = NewChangeInfo(u, 0, m, nil)
	return
}

// NewUpsertInfo creates a new embedded ChangeInfo, with all attributes
// returned on calls to Upsert.
func NewUpsertInfo(u, m int, id interface{}) (c *ChangeInfo) {
	c = NewChangeInfo(u, 0, m, id)
	return
}

// Collection it's an embedded type of mgo.Collection, made to use the
// interfaces it implements, Collectioner, on it's methods signatures.
type Collection struct {
	*mgo.Collection
}

// Find prepares a query using the provided document.  The document may be a
// map or a struct value capable of being marshalled with bson.  The map
// may be a generic one using interface{} for its key and/or values, such as
// bson.M, or it may be a properly typed map.  Providing nil as the document
// is equivalent to providing an empty document such as bson.M{}.
//
// Further details of the query may be tweaked using the resulting Query value,
// and then executed to retrieve results using methods such as One, For,
// Iter, or Tail.
//
// In case the resulting document includes a field named $err or errmsg, which
// are standard ways for MongoDB to return query errors, the returned err will
// be set to a *QueryError value including the Err message and the Code.  In
// those cases, the result argument is still unmarshalled into with the
// received document so that any other custom values may be obtained if
// desired.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Querying
//     http://www.mongodb.org/display/DOCS/Advanced+Queries
//
func (c *Collection) Find(q interface{}) Querier {
	return &Query{c.Collection.Find(q)}
}

// FindId is a convenience helper equivalent to:
//
//     query := collection.Find(bson.M{"_id": id})
//
// See the Find method for more details.
func (c *Collection) FindId(id interface{}) Querier {
	return &Query{c.Collection.FindId(id)}
}

// RemoveAll finds all documents matching the provided selector document
// and removes them from the database.  In case the session is in safe mode
// (see the SetSafe method) and an error happens when attempting the change,
// the returned error will be of type *LastError.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Removing
//
func (c *Collection) RemoveAll(s interface{}) (*ChangeInfo, error) {
	ci, err := c.Collection.RemoveAll(s)
	return &ChangeInfo{ci}, err
}

// UpdateAll finds all documents matching the provided selector document
// and modifies them according to the update document.
// If the session is in safe mode (see SetSafe) details of the executed
// operation are returned in info or an error of type *LastError when
// some problem is detected. It is not an error for the update to not be
// applied on any documents because the selector doesn't match.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
func (c *Collection) UpdateAll(s interface{}, u interface{}) (*ChangeInfo, error) {
	ci, err := c.Collection.UpdateAll(s, u)
	return &ChangeInfo{ci}, err
}

// Upsert finds a single document matching the provided selector document
// and modifies it according to the update document.  If no document matching
// the selector is found, the update document is applied to the selector
// document and the result is inserted in the collection.
// If the session is in safe mode (see SetSafe) details of the executed
// operation are returned in info, or an error of type *LastError when
// some problem is detected.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
func (c *Collection) Upsert(s interface{}, u interface{}) (*ChangeInfo, error) {
	ci, err := c.Collection.Upsert(s, u)
	return &ChangeInfo{ci}, err
}

// UpsertId is a convenience helper equivalent to:
//
//     info, err := collection.Upsert(bson.M{"_id": id}, update)
//
// See the Upsert method for more details.
func (c *Collection) UpsertId(id interface{}, u interface{}) (*ChangeInfo, error) {
	ci, err := c.Collection.UpsertId(id, u)
	return &ChangeInfo{ci}, err
}

// With returns a copy of c that uses session s.
func (c *Collection) With(s *mgo.Session) Collectioner {
	return &Collection{c.Collection.With(s)}
}

// Database it's an embedded type of mgo.Database, made to use the
// interfaces it implements, Databaser, on it's methods signatures.
type Database struct {
	*mgo.Database
}

// C returns a value representing the named collection.
//
// Creating this value is a very lightweight operation, and
// involves no network communication.
func (d *Database) C(n string) Collectioner {
	return &Collection{d.Database.C(n)}
}

// FindRef returns a query that looks for the document in the provided
// reference. If the reference includes the DB field, the document will
// be retrieved from the respective database.
//
// See also the DBRef type and the FindRef method on Session.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Database+References
//
func (d *Database) FindRef(ref *mgo.DBRef) Querier {
	return &Query{d.Database.FindRef(ref)}
}

// With returns a copy of db that uses session s.
func (d *Database) With(s *mgo.Session) Databaser {
	return &Database{d.Database.With(s)}
}

// Query it's an embedded type of mgo.Query, made to use the
// interfaces it implements, Querier, on it's methods signatures.
type Query struct {
	*mgo.Query
}

// The default batch size is defined by the database itself.  As of this
// writing, MongoDB will use an initial size of min(100 docs, 4MB) on the
// first batch, and 4MB on remaining ones.
func (q *Query) Batch(n int) Querier {
	return &Query{q.Query.Batch(n)}
}

// Comment adds a comment to the query to identify it in the database profiler output.
//
// Relevant documentation:
//
//     http://docs.mongodb.org/manual/reference/operator/meta/comment
//     http://docs.mongodb.org/manual/reference/command/profile
//     http://docs.mongodb.org/manual/administration/analyzing-mongodb-performance/#database-profiling
//
func (q *Query) Comment(c string) Querier {
	return &Query{q.Query.Comment(c)}
}

// Hint will include an explicit "hint" in the query to force the server
// to use a specified index, potentially improving performance in some
// situations.  The provided parameters are the fields that compose the
// key of the index to be used.  For details on how the indexKey may be
// built, see the EnsureIndex method.
//
// For example:
//
//     query := collection.Find(bson.M{"firstname": "Joe", "lastname": "Winter"})
//     query.Hint("lastname", "firstname")
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Optimization
//     http://www.mongodb.org/display/DOCS/Query+Optimizer
//
func (q *Query) Hint(idk ...string) Querier {
	return &Query{q.Query.Hint(idk...)}
}

// Limit restricts the maximum number of documents retrieved to n, and also
// changes the batch size to the same value.  Once n documents have been
// returned by Next, the following call will return ErrNotFound.
func (q *Query) Limit(n int) Querier {
	return &Query{q.Query.Limit(n)}
}

// LogReplay enables an option that optimizes queries that are typically
// made on the MongoDB oplog for replaying it. This is an internal
// implementation aspect and most likely uninteresting for other uses.
// It has seen at least one use case, though, so it's exposed via the API.
func (q *Query) LogReplay() Querier {
	return &Query{q.Query.LogReplay()}
}

// Prefetch sets the point at which the next batch of results will be requested.
// When there are p*batch_size remaining documents cached in an Iter, the next
// batch will be requested in background. For instance, when using this:
//
//     query.Batch(200).Prefetch(0.25)
//
// and there are only 50 documents cached in the Iter to be processed, the
// next batch of 200 will be requested. It's possible to change this setting on
// a per-session basis as well, using the SetPrefetch method of Session.
//
// The default prefetch value is 0.25.
func (q *Query) Prefetch(p float64) Querier {
	return &Query{q.Query.Prefetch(p)}
}

// Select enables selecting which fields should be retrieved for the results
// found. For example, the following query would only retrieve the name field:
//
//     err := collection.Find(nil).Select(bson.M{"name": 1}).One(&result)
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Retrieving+a+Subset+of+Fields
//
func (q *Query) Select(sel interface{}) Querier {
	return &Query{q.Query.Select(sel)}
}

// SetMaxScan constrains the query to stop after scanning the specified
// number of documents.
//
// This modifier is generally used to prevent potentially long running
// queries from disrupting performance by scanning through too much data.
func (q *Query) SetMaxScan(n int) Querier {
	return &Query{q.Query.SetMaxScan(n)}
}

// SetMaxTime constrains the query to stop after running for the specified time.
//
// When the time limit is reached MongoDB automatically cancels the query.
// This can be used to efficiently prevent and identify unexpectedly slow queries.
//
// A few important notes about the mechanism enforcing this limit:
//
//  - Requests can block behind locking operations on the server, and that blocking
//    time is not accounted for. In other words, the timer starts ticking only after
//    the actual start of the query when it initially acquires the appropriate lock;
//
//  - Operations are interrupted only at interrupt points where an operation can be
//    safely aborted – the total execution time may exceed the specified value;
//
//  - The limit can be applied to both CRUD operations and commands, but not all
//    commands are interruptible;
//
//  - While iterating over results, computing follow up batches is included in the
//    total time and the iteration continues until the alloted time is over, but
//    network roundtrips are not taken into account for the limit.
//
//  - This limit does not override the inactive cursor timeout for idle cursors
//    (default is 10 min).
//
// This mechanism was introduced in MongoDB 2.6.
//
// Relevant documentation:
//
//   http://blog.mongodb.org/post/83621787773/maxtimems-and-query-optimizer-introspection-in
//
func (q *Query) SetMaxTime(d time.Duration) Querier {
	return &Query{q.Query.SetMaxTime(d)}
}

// Skip skips over the n initial documents from the query results.  Note that
// this only makes sense with capped collections where documents are naturally
// ordered by insertion time, or with sorted results.
func (q *Query) Skip(n int) Querier {
	return &Query{q.Query.Skip(n)}
}

// Snapshot will force the performed query to make use of an available
// index on the _id field to prevent the same document from being returned
// more than once in a single iteration. This might happen without this
// setting in situations when the document changes in size and thus has to
// be moved while the iteration is running.
//
// Because snapshot mode traverses the _id index, it may not be used with
// sorting or explicit hints. It also cannot use any other index for the
// query.
//
// Even with snapshot mode, items inserted or deleted during the query may
// or may not be returned; that is, this mode is not a true point-in-time
// snapshot.
//
// The same effect of Snapshot may be obtained by using any unique index on
// field(s) that will not be modified (best to use Hint explicitly too).
// A non-unique index (such as creation time) may be made unique by
// appending _id to the index when creating it.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/How+to+do+Snapshotted+Queries+in+the+Mongo+Database
//
func (q *Query) Snapshot() Querier {
	return &Query{q.Query.Snapshot()}
}

// Sort asks the database to order returned documents according to the
// provided field names. A field name may be prefixed by - (minus) for
// it to be sorted in reverse order.
//
// For example:
//
//     query1 := collection.Find(nil).Sort("firstname", "lastname")
//     query2 := collection.Find(nil).Sort("-age")
//     query3 := collection.Find(nil).Sort("$natural")
//     query4 := collection.Find(nil).Select(bson.M{"score": bson.M{"$meta": "textScore"}}).Sort("$textScore:score")
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Sorting+and+Natural+Order
//
func (q *Query) Sort(fields ...string) Querier {
	return &Query{q.Query.Sort(fields...)}
}
